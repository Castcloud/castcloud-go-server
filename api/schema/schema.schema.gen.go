package schema

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

func (d *User) Size() (s uint64) {

	{
		l := uint64(len(d.Username))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Password))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Clients))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Clients {

			{
				if d.Clients[k0] != nil {

					{
						s += (*d.Clients[k0]).Size()
					}
					s += 0
				}
			}

			s += 1

		}

	}
	{
		l := uint64(len(d.Subscriptions))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		s += 8 * l

	}
	s += 8
	return
}
func (d *User) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		*(*uint64)(unsafe.Pointer(&buf[0])) = d.ID

	}
	{
		l := uint64(len(d.Username))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.Username)
		i += l
	}
	{
		l := uint64(len(d.Password))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.Password)
		i += l
	}
	{
		l := uint64(len(d.Clients))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		for k0 := range d.Clients {

			{
				if d.Clients[k0] == nil {
					buf[i+8] = 0
				} else {
					buf[i+8] = 1

					{
						nbuf, err := (*d.Clients[k0]).Marshal(buf[i+9:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}
					i += 0
				}
			}

			i += 1

		}
	}
	{
		l := uint64(len(d.Subscriptions))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		for k0 := range d.Subscriptions {

			{

				*(*uint64)(unsafe.Pointer(&buf[i+8])) = d.Subscriptions[k0]

			}

			i += 8

		}
	}
	return buf[:i+8], nil
}

func (d *User) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.ID = *(*uint64)(unsafe.Pointer(&buf[i+0]))

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Username = string(buf[i+8 : i+8+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Password = string(buf[i+8 : i+8+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Clients)) >= l {
			d.Clients = d.Clients[:l]
		} else {
			d.Clients = make([]*Client, l)
		}
		for k0 := range d.Clients {

			{
				if buf[i+8] == 1 {
					if d.Clients[k0] == nil {
						d.Clients[k0] = new(Client)
					}

					{
						ni, err := (*d.Clients[k0]).Unmarshal(buf[i+9:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					i += 0
				} else {
					d.Clients[k0] = nil
				}
			}

			i += 1

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Subscriptions)) >= l {
			d.Subscriptions = d.Subscriptions[:l]
		} else {
			d.Subscriptions = make([]uint64, l)
		}
		for k0 := range d.Subscriptions {

			{

				d.Subscriptions[k0] = *(*uint64)(unsafe.Pointer(&buf[i+8]))

			}

			i += 8

		}
	}
	return i + 8, nil
}

func (d *Client) Size() (s uint64) {

	{
		l := uint64(len(d.Token))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.UUID))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	return
}
func (d *Client) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Token))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Token)
		i += l
	}
	{
		l := uint64(len(d.UUID))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.UUID)
		i += l
	}
	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Name)
		i += l
	}
	return buf[:i+0], nil
}

func (d *Client) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Token = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.UUID = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i+0 : i+0+l])
		i += l
	}
	return i + 0, nil
}

func (d *Setting) Size() (s uint64) {

	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Value))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.ClientUUID))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 9
	return
}
func (d *Setting) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		*(*uint64)(unsafe.Pointer(&buf[0])) = d.ID

	}
	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.Name)
		i += l
	}
	{
		l := uint64(len(d.Value))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.Value)
		i += l
	}
	{
		if d.ClientSpecific {
			buf[i+8] = 1
		} else {
			buf[i+8] = 0
		}
	}
	{
		l := uint64(len(d.ClientUUID))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+9] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+9] = byte(t)
			i++

		}
		copy(buf[i+9:], d.ClientUUID)
		i += l
	}
	return buf[:i+9], nil
}

func (d *Setting) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.ID = *(*uint64)(unsafe.Pointer(&buf[i+0]))

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i+8 : i+8+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Value = string(buf[i+8 : i+8+l])
		i += l
	}
	{
		d.ClientSpecific = buf[i+8] == 1
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+9] & 0x7F)
			for buf[i+9]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+9]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.ClientUUID = string(buf[i+9 : i+9+l])
		i += l
	}
	return i + 9, nil
}

func (d *Cast) Size() (s uint64) {

	{
		l := uint64(len(d.URL))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.FeedBytes))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 8
	return
}
func (d *Cast) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		*(*uint64)(unsafe.Pointer(&buf[0])) = d.ID

	}
	{
		l := uint64(len(d.URL))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.URL)
		i += l
	}
	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.Name)
		i += l
	}
	{
		l := uint64(len(d.FeedBytes))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.FeedBytes)
		i += l
	}
	return buf[:i+8], nil
}

func (d *Cast) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.ID = *(*uint64)(unsafe.Pointer(&buf[i+0]))

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.URL = string(buf[i+8 : i+8+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i+8 : i+8+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.FeedBytes)) >= l {
			d.FeedBytes = d.FeedBytes[:l]
		} else {
			d.FeedBytes = make([]byte, l)
		}
		copy(d.FeedBytes, buf[i+8:])
		i += l
	}
	return i + 8, nil
}

func (d *Episode) Size() (s uint64) {

	{
		if d.LastEvent != nil {

			{
				s += (*d.LastEvent).Size()
			}
			s += 0
		}
	}
	{
		l := uint64(len(d.FeedBytes))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.GUID))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 25
	return
}
func (d *Episode) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		*(*uint64)(unsafe.Pointer(&buf[0])) = d.ID

	}
	{

		*(*uint64)(unsafe.Pointer(&buf[8])) = d.CastID

	}
	{
		if d.LastEvent == nil {
			buf[i+16] = 0
		} else {
			buf[i+16] = 1

			{
				nbuf, err := (*d.LastEvent).Marshal(buf[17:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}
			i += 0
		}
	}
	{
		l := uint64(len(d.FeedBytes))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+17] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+17] = byte(t)
			i++

		}
		copy(buf[i+17:], d.FeedBytes)
		i += l
	}
	{
		l := uint64(len(d.GUID))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+17] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+17] = byte(t)
			i++

		}
		copy(buf[i+17:], d.GUID)
		i += l
	}
	{

		*(*int64)(unsafe.Pointer(&buf[i+17])) = d.CrawlTS

	}
	return buf[:i+25], nil
}

func (d *Episode) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.ID = *(*uint64)(unsafe.Pointer(&buf[i+0]))

	}
	{

		d.CastID = *(*uint64)(unsafe.Pointer(&buf[i+8]))

	}
	{
		if buf[i+16] == 1 {
			if d.LastEvent == nil {
				d.LastEvent = new(Event)
			}

			{
				ni, err := (*d.LastEvent).Unmarshal(buf[i+17:])
				if err != nil {
					return 0, err
				}
				i += ni
			}
			i += 0
		} else {
			d.LastEvent = nil
		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+17] & 0x7F)
			for buf[i+17]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+17]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.FeedBytes)) >= l {
			d.FeedBytes = d.FeedBytes[:l]
		} else {
			d.FeedBytes = make([]byte, l)
		}
		copy(d.FeedBytes, buf[i+17:])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+17] & 0x7F)
			for buf[i+17]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+17]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.GUID = string(buf[i+17 : i+17+l])
		i += l
	}
	{

		d.CrawlTS = *(*int64)(unsafe.Pointer(&buf[i+17]))

	}
	return i + 25, nil
}

func (d *Event) Size() (s uint64) {

	{
		l := uint64(len(d.ClientName))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.ClientDescription))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.ClientUUID))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 28
	return
}
func (d *Event) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		*(*int32)(unsafe.Pointer(&buf[0])) = d.Type

	}
	{

		*(*uint64)(unsafe.Pointer(&buf[4])) = d.EpisodeID

	}
	{

		*(*int32)(unsafe.Pointer(&buf[12])) = d.PositionTS

	}
	{

		*(*uint64)(unsafe.Pointer(&buf[16])) = d.ClientTS

	}
	{

		*(*int32)(unsafe.Pointer(&buf[24])) = d.ConcurrentOrder

	}
	{
		l := uint64(len(d.ClientName))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+28] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+28] = byte(t)
			i++

		}
		copy(buf[i+28:], d.ClientName)
		i += l
	}
	{
		l := uint64(len(d.ClientDescription))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+28] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+28] = byte(t)
			i++

		}
		copy(buf[i+28:], d.ClientDescription)
		i += l
	}
	{
		l := uint64(len(d.ClientUUID))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+28] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+28] = byte(t)
			i++

		}
		copy(buf[i+28:], d.ClientUUID)
		i += l
	}
	return buf[:i+28], nil
}

func (d *Event) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Type = *(*int32)(unsafe.Pointer(&buf[i+0]))

	}
	{

		d.EpisodeID = *(*uint64)(unsafe.Pointer(&buf[i+4]))

	}
	{

		d.PositionTS = *(*int32)(unsafe.Pointer(&buf[i+12]))

	}
	{

		d.ClientTS = *(*uint64)(unsafe.Pointer(&buf[i+16]))

	}
	{

		d.ConcurrentOrder = *(*int32)(unsafe.Pointer(&buf[i+24]))

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+28] & 0x7F)
			for buf[i+28]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+28]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.ClientName = string(buf[i+28 : i+28+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+28] & 0x7F)
			for buf[i+28]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+28]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.ClientDescription = string(buf[i+28 : i+28+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+28] & 0x7F)
			for buf[i+28]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+28]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.ClientUUID = string(buf[i+28 : i+28+l])
		i += l
	}
	return i + 28, nil
}

func (d *Label) Size() (s uint64) {

	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Content))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 10
	return
}
func (d *Label) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		*(*uint64)(unsafe.Pointer(&buf[0])) = d.ID

	}
	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.Name)
		i += l
	}
	{
		l := uint64(len(d.Content))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.Content)
		i += l
	}
	{
		if d.Expanded {
			buf[i+8] = 1
		} else {
			buf[i+8] = 0
		}
	}
	{
		if d.Root {
			buf[i+9] = 1
		} else {
			buf[i+9] = 0
		}
	}
	return buf[:i+10], nil
}

func (d *Label) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.ID = *(*uint64)(unsafe.Pointer(&buf[i+0]))

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i+8 : i+8+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Content = string(buf[i+8 : i+8+l])
		i += l
	}
	{
		d.Expanded = buf[i+8] == 1
	}
	{
		d.Root = buf[i+9] == 1
	}
	return i + 10, nil
}
