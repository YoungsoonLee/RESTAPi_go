package models

import "time"

/**
  payment_items                //유료, 무료 관련 코인 아이템 테이블. 실제 결제시 참조 되는 테이블로 매우 중요한 테이블 이다.
     item_id                      // unique, auto increase
     category_cid                 // paymentCategory 테이블의 pk인 cid
     item_name                    //
     item_description             //
     pg_id                        // payment_category.category_cid 100번대 일 경우 셋팅 됨. payment_gateway.pg_id
     currency                     // default: 'USD'..defaultTo('USD')
     price                        // payment_category.category_id 100번대 일 경우 셋팅 됨. 나머진 0
     amount                       // 실제 적립되는 cyber coin 양
     created_at                   // .defaultTo(knex.fn.now()
     updated_at
     closed_at

     * discount는 보너스 어마운트를 부여 하는 방식.
*/
type PaymentItem struct {
	ItemID          int       `orm:"column(ItemID);pk;auto" json:"itemid"`         // unique, auto increase
	CategoryID      int       `orm:"column(CategoryID);" json:"categoryid"`        // not null, // paymentCategory 테이블의 pk인 cid
	ItemName        string    `orm:"size(1000);" json:"item_name"`                 // not null,
	ItemDescription string    `orm:"size(2000);" json:"item_description"`          // not null,
	PgID            int       `orm:"column(PgID);null" json:"pgid"`                // payment_category.category_cid 100번대(유료) 일 경우 payment_gateway.pg_id가 셋팅 됨.
	Currency        string    `orm:"size(3);default(USD)" json:"currency"`         // not null, default 'USD'
	Price           int       `orm:"default(0)" json:"price"`                      // not null, payment_category.category_id 100번대 일 경우 셋팅 됨. 나머진 0
	Amount          int       `orm:"default(0)" json:"amount"`                     // 실제 적립되는 cyber coin 양
	CreateAt        time.Time `orm:"type(datetime);auto_now_add" json:"create_at"` // first save
	UpdateAt        time.Time `orm:"type(datetime);auto_now" json:"update_at"`     // eveytime save
	CloseAt         time.Time `orm:"type(datetime);null" json:"close_at"`          //
}
