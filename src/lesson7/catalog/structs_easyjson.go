// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package catalog

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

func easyjson6a975c40DecodeLesson7Catalog(in *jlexer.Lexer, out *Product) {
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
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a975c40EncodeLesson7Catalog(out *jwriter.Writer, in Product) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Product) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeLesson7Catalog(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Product) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeLesson7Catalog(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Product) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeLesson7Catalog(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Product) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeLesson7Catalog(l, v)
}
func easyjson6a975c40DecodeLesson7Catalog1(in *jlexer.Lexer, out *OrderItem) {
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
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a975c40EncodeLesson7Catalog1(out *jwriter.Writer, in OrderItem) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OrderItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeLesson7Catalog1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeLesson7Catalog1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeLesson7Catalog1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeLesson7Catalog1(l, v)
}
func easyjson6a975c40DecodeLesson7Catalog2(in *jlexer.Lexer, out *OrderDeliveryData) {
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
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a975c40EncodeLesson7Catalog2(out *jwriter.Writer, in OrderDeliveryData) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OrderDeliveryData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeLesson7Catalog2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderDeliveryData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeLesson7Catalog2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderDeliveryData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeLesson7Catalog2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderDeliveryData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeLesson7Catalog2(l, v)
}
func easyjson6a975c40DecodeLesson7Catalog3(in *jlexer.Lexer, out *OrderCustomerData) {
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
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a975c40EncodeLesson7Catalog3(out *jwriter.Writer, in OrderCustomerData) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OrderCustomerData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeLesson7Catalog3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderCustomerData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeLesson7Catalog3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderCustomerData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeLesson7Catalog3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderCustomerData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeLesson7Catalog3(l, v)
}
func easyjson6a975c40DecodeLesson7Catalog4(in *jlexer.Lexer, out *Order) {
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
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a975c40EncodeLesson7Catalog4(out *jwriter.Writer, in Order) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Order) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeLesson7Catalog4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Order) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeLesson7Catalog4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Order) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeLesson7Catalog4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Order) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeLesson7Catalog4(l, v)
}
func easyjson6a975c40DecodeLesson7Catalog5(in *jlexer.Lexer, out *Category) {
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
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6a975c40EncodeLesson7Catalog5(out *jwriter.Writer, in Category) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Category) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeLesson7Catalog5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Category) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeLesson7Catalog5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Category) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeLesson7Catalog5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Category) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeLesson7Catalog5(l, v)
}
