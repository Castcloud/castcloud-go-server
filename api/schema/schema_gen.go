package schema

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Client) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Token":
			z.Token, err = dc.ReadString()
			if err != nil {
				return
			}
		case "UUID":
			z.UUID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Client) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "Token"
	err = en.Append(0x83, 0xa5, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Token)
	if err != nil {
		return
	}
	// write "UUID"
	err = en.Append(0xa4, 0x55, 0x55, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.UUID)
	if err != nil {
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Client) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Token"
	o = append(o, 0x83, 0xa5, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
	o = msgp.AppendString(o, z.Token)
	// string "UUID"
	o = append(o, 0xa4, 0x55, 0x55, 0x49, 0x44)
	o = msgp.AppendString(o, z.UUID)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Client) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Token":
			z.Token, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "UUID":
			z.UUID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z Client) Msgsize() (s int) {
	s = 1 + 6 + msgp.StringPrefixSize + len(z.Token) + 5 + msgp.StringPrefixSize + len(z.UUID) + 5 + msgp.StringPrefixSize + len(z.Name)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Cast) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "URL":
			z.URL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FeedMsgp":
			z.FeedMsgp, err = dc.ReadBytes(z.FeedMsgp)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Cast) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "ID"
	err = en.Append(0x84, 0xa2, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.ID)
	if err != nil {
		return
	}
	// write "URL"
	err = en.Append(0xa3, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.URL)
	if err != nil {
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "FeedMsgp"
	err = en.Append(0xa8, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.FeedMsgp)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Cast) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "ID"
	o = append(o, 0x84, 0xa2, 0x49, 0x44)
	o = msgp.AppendUint64(o, z.ID)
	// string "URL"
	o = append(o, 0xa3, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.URL)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "FeedMsgp"
	o = append(o, 0xa8, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x70)
	o = msgp.AppendBytes(o, z.FeedMsgp)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Cast) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "URL":
			z.URL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FeedMsgp":
			z.FeedMsgp, bts, err = msgp.ReadBytesBytes(bts, z.FeedMsgp)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Cast) Msgsize() (s int) {
	s = 1 + 3 + msgp.Uint64Size + 4 + msgp.StringPrefixSize + len(z.URL) + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.BytesPrefixSize + len(z.FeedMsgp)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Episode) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "CastID":
			z.CastID, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "LastEvent":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					return
				}
				z.LastEvent = nil
			} else {
				if z.LastEvent == nil {
					z.LastEvent = new(Event)
				}
				var isz uint32
				isz, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				for isz > 0 {
					isz--
					field, err = dc.ReadMapKeyPtr()
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "ID":
						z.LastEvent.ID, err = dc.ReadUint64()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
			}
		case "FeedMsgp":
			z.FeedMsgp, err = dc.ReadBytes(z.FeedMsgp)
			if err != nil {
				return
			}
		case "GUID":
			z.GUID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "CrawlTS":
			z.CrawlTS, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Episode) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "ID"
	err = en.Append(0x86, 0xa2, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.ID)
	if err != nil {
		return
	}
	// write "CastID"
	err = en.Append(0xa6, 0x43, 0x61, 0x73, 0x74, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.CastID)
	if err != nil {
		return
	}
	// write "LastEvent"
	err = en.Append(0xa9, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	if z.LastEvent == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		// map header, size 1
		// write "ID"
		err = en.Append(0x81, 0xa2, 0x49, 0x44)
		if err != nil {
			return err
		}
		err = en.WriteUint64(z.LastEvent.ID)
		if err != nil {
			return
		}
	}
	// write "FeedMsgp"
	err = en.Append(0xa8, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x70)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.FeedMsgp)
	if err != nil {
		return
	}
	// write "GUID"
	err = en.Append(0xa4, 0x47, 0x55, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.GUID)
	if err != nil {
		return
	}
	// write "CrawlTS"
	err = en.Append(0xa7, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x54, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.CrawlTS)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Episode) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "ID"
	o = append(o, 0x86, 0xa2, 0x49, 0x44)
	o = msgp.AppendUint64(o, z.ID)
	// string "CastID"
	o = append(o, 0xa6, 0x43, 0x61, 0x73, 0x74, 0x49, 0x44)
	o = msgp.AppendUint64(o, z.CastID)
	// string "LastEvent"
	o = append(o, 0xa9, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74)
	if z.LastEvent == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 1
		// string "ID"
		o = append(o, 0x81, 0xa2, 0x49, 0x44)
		o = msgp.AppendUint64(o, z.LastEvent.ID)
	}
	// string "FeedMsgp"
	o = append(o, 0xa8, 0x46, 0x65, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x70)
	o = msgp.AppendBytes(o, z.FeedMsgp)
	// string "GUID"
	o = append(o, 0xa4, 0x47, 0x55, 0x49, 0x44)
	o = msgp.AppendString(o, z.GUID)
	// string "CrawlTS"
	o = append(o, 0xa7, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x54, 0x53)
	o = msgp.AppendInt64(o, z.CrawlTS)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Episode) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "CastID":
			z.CastID, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "LastEvent":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.LastEvent = nil
			} else {
				if z.LastEvent == nil {
					z.LastEvent = new(Event)
				}
				var isz uint32
				isz, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for isz > 0 {
					isz--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "ID":
						z.LastEvent.ID, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							return
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "FeedMsgp":
			z.FeedMsgp, bts, err = msgp.ReadBytesBytes(bts, z.FeedMsgp)
			if err != nil {
				return
			}
		case "GUID":
			z.GUID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "CrawlTS":
			z.CrawlTS, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *Episode) Msgsize() (s int) {
	s = 1 + 3 + msgp.Uint64Size + 7 + msgp.Uint64Size + 10
	if z.LastEvent == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 3 + msgp.Uint64Size
	}
	s += 9 + msgp.BytesPrefixSize + len(z.FeedMsgp) + 5 + msgp.StringPrefixSize + len(z.GUID) + 8 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Event) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, err = dc.ReadUint64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Event) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "ID"
	err = en.Append(0x81, 0xa2, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.ID)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Event) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "ID"
	o = append(o, 0x81, 0xa2, 0x49, 0x44)
	o = msgp.AppendUint64(o, z.ID)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Event) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z Event) Msgsize() (s int) {
	s = 1 + 3 + msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *User) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "Username":
			z.Username, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Password":
			z.Password, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Clients":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Clients) >= int(xsz) {
				z.Clients = z.Clients[:xsz]
			} else {
				z.Clients = make([]*Client, xsz)
			}
			for xvk := range z.Clients {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Clients[xvk] = nil
				} else {
					if z.Clients[xvk] == nil {
						z.Clients[xvk] = new(Client)
					}
					var isz uint32
					isz, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for isz > 0 {
						isz--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "Token":
							z.Clients[xvk].Token, err = dc.ReadString()
							if err != nil {
								return
							}
						case "UUID":
							z.Clients[xvk].UUID, err = dc.ReadString()
							if err != nil {
								return
							}
						case "Name":
							z.Clients[xvk].Name, err = dc.ReadString()
							if err != nil {
								return
							}
						default:
							err = dc.Skip()
							if err != nil {
								return
							}
						}
					}
				}
			}
		case "Subscriptions":
			var xsz uint32
			xsz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Subscriptions) >= int(xsz) {
				z.Subscriptions = z.Subscriptions[:xsz]
			} else {
				z.Subscriptions = make([]uint64, xsz)
			}
			for bzg := range z.Subscriptions {
				z.Subscriptions[bzg], err = dc.ReadUint64()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *User) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "ID"
	err = en.Append(0x85, 0xa2, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.ID)
	if err != nil {
		return
	}
	// write "Username"
	err = en.Append(0xa8, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Username)
	if err != nil {
		return
	}
	// write "Password"
	err = en.Append(0xa8, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Password)
	if err != nil {
		return
	}
	// write "Clients"
	err = en.Append(0xa7, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Clients)))
	if err != nil {
		return
	}
	for xvk := range z.Clients {
		if z.Clients[xvk] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			// map header, size 3
			// write "Token"
			err = en.Append(0x83, 0xa5, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Clients[xvk].Token)
			if err != nil {
				return
			}
			// write "UUID"
			err = en.Append(0xa4, 0x55, 0x55, 0x49, 0x44)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Clients[xvk].UUID)
			if err != nil {
				return
			}
			// write "Name"
			err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
			if err != nil {
				return err
			}
			err = en.WriteString(z.Clients[xvk].Name)
			if err != nil {
				return
			}
		}
	}
	// write "Subscriptions"
	err = en.Append(0xad, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Subscriptions)))
	if err != nil {
		return
	}
	for bzg := range z.Subscriptions {
		err = en.WriteUint64(z.Subscriptions[bzg])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *User) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "ID"
	o = append(o, 0x85, 0xa2, 0x49, 0x44)
	o = msgp.AppendUint64(o, z.ID)
	// string "Username"
	o = append(o, 0xa8, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Username)
	// string "Password"
	o = append(o, 0xa8, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64)
	o = msgp.AppendString(o, z.Password)
	// string "Clients"
	o = append(o, 0xa7, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Clients)))
	for xvk := range z.Clients {
		if z.Clients[xvk] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 3
			// string "Token"
			o = append(o, 0x83, 0xa5, 0x54, 0x6f, 0x6b, 0x65, 0x6e)
			o = msgp.AppendString(o, z.Clients[xvk].Token)
			// string "UUID"
			o = append(o, 0xa4, 0x55, 0x55, 0x49, 0x44)
			o = msgp.AppendString(o, z.Clients[xvk].UUID)
			// string "Name"
			o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.Clients[xvk].Name)
		}
	}
	// string "Subscriptions"
	o = append(o, 0xad, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Subscriptions)))
	for bzg := range z.Subscriptions {
		o = msgp.AppendUint64(o, z.Subscriptions[bzg])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *User) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for isz > 0 {
		isz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			z.ID, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "Username":
			z.Username, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Password":
			z.Password, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Clients":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Clients) >= int(xsz) {
				z.Clients = z.Clients[:xsz]
			} else {
				z.Clients = make([]*Client, xsz)
			}
			for xvk := range z.Clients {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Clients[xvk] = nil
				} else {
					if z.Clients[xvk] == nil {
						z.Clients[xvk] = new(Client)
					}
					var isz uint32
					isz, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for isz > 0 {
						isz--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "Token":
							z.Clients[xvk].Token, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "UUID":
							z.Clients[xvk].UUID, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "Name":
							z.Clients[xvk].Name, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
				}
			}
		case "Subscriptions":
			var xsz uint32
			xsz, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Subscriptions) >= int(xsz) {
				z.Subscriptions = z.Subscriptions[:xsz]
			} else {
				z.Subscriptions = make([]uint64, xsz)
			}
			for bzg := range z.Subscriptions {
				z.Subscriptions[bzg], bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

func (z *User) Msgsize() (s int) {
	s = 1 + 3 + msgp.Uint64Size + 9 + msgp.StringPrefixSize + len(z.Username) + 9 + msgp.StringPrefixSize + len(z.Password) + 8 + msgp.ArrayHeaderSize
	for xvk := range z.Clients {
		if z.Clients[xvk] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 6 + msgp.StringPrefixSize + len(z.Clients[xvk].Token) + 5 + msgp.StringPrefixSize + len(z.Clients[xvk].UUID) + 5 + msgp.StringPrefixSize + len(z.Clients[xvk].Name)
		}
	}
	s += 14 + msgp.ArrayHeaderSize + (len(z.Subscriptions) * (msgp.Uint64Size))
	return
}
