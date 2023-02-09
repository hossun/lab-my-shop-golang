package inventory

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
	"fmt"
	"inventory/external"
)

type Inventory struct {
	gorm.Model
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	Stock int `json:"stock"`

}

func (self *Inventory) onPostPersist() (err error){

	return nil
}
func (self *Inventory) onPrePersist() (err error){ return nil }
func (self *Inventory) onPreUpdate() (err error){ return nil }
func (self *Inventory) onPostUpdate() (err error){ return nil }
func (self *Inventory) onPreRemove() (err error){ return nil }
func (self *Inventory) onPostRemove() (err error){ return nil }


func DecreaseStock(orderPlaced *OrderPlaced){
	/** Example 1:  new item
	inventory := &Inventory{}
	inventoryrepository.save(inventory)

	*/

	/** Example 2:  finding and process
	id, _ := strconv.ParseInt(orderPlaced.id, 10, 64)
	inventory, err := inventoryrepository().FindById(int(id))
	if err != nil {

	}
	*/
}
func IncreaseStock(orderCancelled *OrderCancelled){
	/** Example 1:  new item
	inventory := &Inventory{}
	inventoryrepository.save(inventory)

	*/

	/** Example 2:  finding and process
	id, _ := strconv.ParseInt(orderCancelled.id, 10, 64)
	inventory, err := inventoryrepository().FindById(int(id))
	if err != nil {

	}
	*/
}
