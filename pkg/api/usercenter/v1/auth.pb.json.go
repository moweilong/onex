// This file defines the Protobuf messages for managing Users.
//

// Code generated by protoc-gen-go-json. DO NOT EDIT.
// source: usercenter/v1/auth.proto

package v1

import (
	"google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON implements json.Marshaler
func (msg *AuthenticateRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *AuthenticateRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *AuthenticateResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *AuthenticateResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *AuthorizeRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *AuthorizeRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *AuthorizeResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *AuthorizeResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *AuthRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *AuthRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *AuthResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *AuthResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{}.Unmarshal(b, msg)
}
