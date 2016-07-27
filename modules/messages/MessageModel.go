package messages

type Message struct {
  ID        int     `json:"id"`
  Author    string  `json:"author"`
  Message   string  `json:"message"`
}

func (m *Message) GetAuthor() string {
  return m.Author
}

func (m *Message) GetMessage() string {
  return m.Message
}
