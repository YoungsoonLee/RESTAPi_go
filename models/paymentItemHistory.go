package models

/**
  payment_items_history         //paymentItems 이력 관리 테이블.(관리툴 이용시)
     id                          // unique, auto increase, pk
     item_id                     //
     category_cid                // paymentCategory 테이블의 cid
     item_name                   //
     item_description            //
     pg_id                       // paymentCategory.categoryId가 100번대 일 경우 셋팅 됨
     currency                    // default: 'USD'.
     price                       // paymentCategory.categoryId가 100번대 일 경우 셋팅 됨
     amount                      // 실제 적립되는 cyber coin 양
     updated_at
     admin                       // 누가 변경을 했는지. ex) 관리툴에서...
*/
/*
exports.up = function(knex, Promise) {
    return Promise.all([
        knex.schema.createTable('payment_items_history', function(table) {
            table
                .increments('id')
                .unsigned()
                .primary();
            table
                .integer('item_id')
                .unsigned()
                .notNullable();
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
            table.timestamp('updated_at').defaultTo(knex.fn.now());
            table.string('admin');
        })
    ]);
};
*/
