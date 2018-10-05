package models

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
/*
exports.up = function(knex, Promise) {
    return Promise.all([
        knex.schema.createTable('payment_items', function(table) {
            table
                .increments('item_id')
                .unsigned()
                .primary();
            table
                .integer('category_cid')
                .unsigned()
                .notNullable();
            table.string('item_name').notNullable();
            table.string('item_description').notNullable();
            table
                .integer('pg_id')
                .unsigned()
                .notNullable();
            table.string('currency', 3).defaultTo('USD');
            table
                .float('price')
                .unsigned()
                .notNullable();
            table
                .integer('amount')
                .unsigned()
                .notNullable();
            table.timestamp('created_at').defaultTo(knex.fn.now());
            table.timestamp('updated_at');
            table.timestamp('closed_at');
        })
    ]);
};
*/
