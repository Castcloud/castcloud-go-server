package schema

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/Castcloud/castcloud-go-server/Godeps/_workspace/src/github.com/tinylib/msgp/msgp"
)

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
		case "Type":
			z.Type, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "EpisodeID":
			z.EpisodeID, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "PositionTS":
			z.PositionTS, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "ClientTS":
			z.ClientTS, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "ConcurrentOrder":
			z.ConcurrentOrder, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "ClientName":
			z.ClientName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ClientDescription":
			z.ClientDescription, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ClientUUID":
			z.ClientUUID, err = dc.ReadString()
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
func (z *Event) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 8
	// write "Type"
	err = en.Append(0x88, 0xa4, 0x54, 0x79, 0x70, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Type)
	if err != nil {
		return
	}
	// write "EpisodeID"
	err = en.Append(0xa9, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.EpisodeID)
	if err != nil {
		return
	}
	// write "PositionTS"
	err = en.Append(0xaa, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.PositionTS)
	if err != nil {
		return
	}
	// write "ClientTS"
	err = en.Append(0xa8, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.ClientTS)
	if err != nil {
		return
	}
	// write "ConcurrentOrder"
	err = en.Append(0xaf, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.ConcurrentOrder)
	if err != nil {
		return
	}
	// write "ClientName"
	err = en.Append(0xaa, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ClientName)
	if err != nil {
		return
	}
	// write "ClientDescription"
	err = en.Append(0xb1, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ClientDescription)
	if err != nil {
		return
	}
	// write "ClientUUID"
	err = en.Append(0xaa, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ClientUUID)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Event) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "Type"
	o = append(o, 0x88, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt(o, z.Type)
	// string "EpisodeID"
	o = append(o, 0xa9, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x49, 0x44)
	o = msgp.AppendUint64(o, z.EpisodeID)
	// string "PositionTS"
	o = append(o, 0xaa, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x53)
	o = msgp.AppendInt(o, z.PositionTS)
	// string "ClientTS"
	o = append(o, 0xa8, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x53)
	o = msgp.AppendUint64(o, z.ClientTS)
	// string "ConcurrentOrder"
	o = append(o, 0xaf, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72)
	o = msgp.AppendInt(o, z.ConcurrentOrder)
	// string "ClientName"
	o = append(o, 0xaa, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.ClientName)
	// string "ClientDescription"
	o = append(o, 0xb1, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.ClientDescription)
	// string "ClientUUID"
	o = append(o, 0xaa, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x55, 0x55, 0x49, 0x44)
	o = msgp.AppendString(o, z.ClientUUID)
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
		case "Type":
			z.Type, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "EpisodeID":
			z.EpisodeID, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "PositionTS":
			z.PositionTS, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "ClientTS":
			z.ClientTS, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "ConcurrentOrder":
			z.ConcurrentOrder, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "ClientName":
			z.ClientName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ClientDescription":
			z.ClientDescription, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ClientUUID":
			z.ClientUUID, bts, err = msgp.ReadStringBytes(bts)
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

func (z *Event) Msgsize() (s int) {
	s = 1 + 5 + msgp.IntSize + 10 + msgp.Uint64Size + 11 + msgp.IntSize + 9 + msgp.Uint64Size + 16 + msgp.IntSize + 11 + msgp.StringPrefixSize + len(z.ClientName) + 18 + msgp.StringPrefixSize + len(z.ClientDescription) + 11 + msgp.StringPrefixSize + len(z.ClientUUID)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Label) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Content":
			z.Content, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Expanded":
			z.Expanded, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "Root":
			z.Root, err = dc.ReadBool()
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
func (z *Label) EncodeMsg(en *msgp.Writer) (err error) {
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
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Content"
	err = en.Append(0xa7, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Content)
	if err != nil {
		return
	}
	// write "Expanded"
	err = en.Append(0xa8, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Expanded)
	if err != nil {
		return
	}
	// write "Root"
	err = en.Append(0xa4, 0x52, 0x6f, 0x6f, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Root)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Label) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "ID"
	o = append(o, 0x85, 0xa2, 0x49, 0x44)
	o = msgp.AppendUint64(o, z.ID)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Content"
	o = append(o, 0xa7, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Content)
	// string "Expanded"
	o = append(o, 0xa8, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64)
	o = msgp.AppendBool(o, z.Expanded)
	// string "Root"
	o = append(o, 0xa4, 0x52, 0x6f, 0x6f, 0x74)
	o = msgp.AppendBool(o, z.Root)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Label) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Content":
			z.Content, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Expanded":
			z.Expanded, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "Root":
			z.Root, bts, err = msgp.ReadBoolBytes(bts)
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

func (z *Label) Msgsize() (s int) {
	s = 1 + 3 + msgp.Uint64Size + 5 + msgp.StringPrefixSize + len(z.Name) + 8 + msgp.StringPrefixSize + len(z.Content) + 9 + msgp.BoolSize + 5 + msgp.BoolSize
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
				err = z.LastEvent.DecodeMsg(dc)
				if err != nil {
					return
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
		err = z.LastEvent.EncodeMsg(en)
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
		o, err = z.LastEvent.MarshalMsg(o)
		if err != nil {
			return
		}
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
				bts, err = z.LastEvent.UnmarshalMsg(bts)
				if err != nil {
					return
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
		s += z.LastEvent.Msgsize()
	}
	s += 9 + msgp.BytesPrefixSize + len(z.FeedMsgp) + 5 + msgp.StringPrefixSize + len(z.GUID) + 8 + msgp.Int64Size
	return
}
