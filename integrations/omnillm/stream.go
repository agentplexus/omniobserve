package omnillm

import (
	"io"
	"strings"

	"github.com/agentplexus/omnillm"
	"github.com/agentplexus/omnillm/provider"

	"github.com/agentplexus/omniobserve/llmops"
)

// observedStream wraps a provider.ChatCompletionStream to capture
// streaming content and record it when the stream ends.
type observedStream struct {
	stream        provider.ChatCompletionStream
	span          llmops.Span
	info          omnillm.LLMCallInfo
	contentBuffer strings.Builder
	ended         bool
}

// Recv receives the next chunk from the stream.
// It buffers the content and ends the span when the stream completes.
func (s *observedStream) Recv() (*provider.ChatCompletionChunk, error) {
	chunk, err := s.stream.Recv()

	if err == io.EOF {
		// Stream complete - finalize span
		s.finalizeSpan(nil)
		return chunk, err
	}

	if err != nil {
		s.finalizeSpan(err)
		return chunk, err
	}

	// Buffer content from chunk
	if chunk != nil && len(chunk.Choices) > 0 && chunk.Choices[0].Delta != nil {
		s.contentBuffer.WriteString(chunk.Choices[0].Delta.Content)
	}

	return chunk, nil
}

// Close closes the underlying stream.
func (s *observedStream) Close() error {
	// Ensure span is ended if Close is called before EOF
	s.finalizeSpan(nil)
	return s.stream.Close()
}

// finalizeSpan ends the span with the buffered content.
func (s *observedStream) finalizeSpan(err error) {
	if s.ended {
		return
	}
	s.ended = true

	if err != nil {
		_ = s.span.End(llmops.WithEndError(err))
		return
	}

	// Set the buffered output
	if s.contentBuffer.Len() > 0 {
		_ = s.span.SetOutput(s.contentBuffer.String())
	}

	_ = s.span.End()
}
