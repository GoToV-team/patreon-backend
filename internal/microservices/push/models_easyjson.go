// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package push

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodePatreonInternalMicroservicesPush(in *jlexer.Lexer, out *PostInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "creator_id":
			out.CreatorId = int64(in.Int64())
		case "post_id":
			out.PostId = int64(in.Int64())
		case "post_title":
			out.PostTitle = string(in.String())
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodePatreonInternalMicroservicesPush(out *jwriter.Writer, in PostInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"creator_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.CreatorId))
	}
	{
		const prefix string = ",\"post_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.PostId))
	}
	{
		const prefix string = ",\"post_title\":"
		out.RawString(prefix)
		out.String(string(in.PostTitle))
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.Raw((in.Date).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodePatreonInternalMicroservicesPush(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodePatreonInternalMicroservicesPush(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodePatreonInternalMicroservicesPush(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodePatreonInternalMicroservicesPush(l, v)
}
func easyjsonD2b7633eDecodePatreonInternalMicroservicesPush1(in *jlexer.Lexer, out *PaymentApply) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "token":
			out.Token = string(in.String())
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodePatreonInternalMicroservicesPush1(out *jwriter.Writer, in PaymentApply) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"token\":"
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.Raw((in.Date).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PaymentApply) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodePatreonInternalMicroservicesPush1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PaymentApply) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodePatreonInternalMicroservicesPush1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PaymentApply) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodePatreonInternalMicroservicesPush1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PaymentApply) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodePatreonInternalMicroservicesPush1(l, v)
}
func easyjsonD2b7633eDecodePatreonInternalMicroservicesPush2(in *jlexer.Lexer, out *CommentInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "comment_id":
			out.CommentId = int64(in.Int64())
		case "author_id":
			out.AuthorId = int64(in.Int64())
		case "post_id":
			out.PostId = int64(in.Int64())
		case "date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Date).UnmarshalJSON(data))
			}
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodePatreonInternalMicroservicesPush2(out *jwriter.Writer, in CommentInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"comment_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.CommentId))
	}
	{
		const prefix string = ",\"author_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.AuthorId))
	}
	{
		const prefix string = ",\"post_id\":"
		out.RawString(prefix)
		out.Int64(int64(in.PostId))
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.Raw((in.Date).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CommentInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodePatreonInternalMicroservicesPush2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CommentInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodePatreonInternalMicroservicesPush2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CommentInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodePatreonInternalMicroservicesPush2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CommentInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodePatreonInternalMicroservicesPush2(l, v)
}
