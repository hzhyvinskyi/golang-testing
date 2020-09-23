package cases

import "fmt"

type CaseOwner struct {
	ID string
}

func (c *CaseOwner) getID() string {
	return c.ID
}

func (c *CaseOwner) update(p *Payload) {
	fmt.Printf("Trigger update logic for %s.\nEvent ID: %d\nTimestamp: %s", p.Command, p.ID, p.Timestamp)
}
