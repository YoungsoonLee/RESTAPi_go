package models

/**
  payment_category            //아레 item이 어느 카데고리인지 관리 하는 테이블. 유료 무료, 무료일 경우 어떠한 무료 인지 구분 짓는 테이블 이다.(주로 통계용이 주목적이다.)
     cid                         // unique, auto increse. pk
     category_id                 // 100: 유료 충전용, 200: 무료 rewars, 300: 무료 bonus
     category_description        //
     sub_category_id             // 0:paid(charge paid coin), ex)category_id를 다시 상세화 할 때 사용. 주로 통계용
     sub_category_description    //
     created_at                  // .defaultTo(knex.fn.now())
     closed_at                   //         [description]
*/
/*
exports.up = function(knex, Promise) {
    return Promise.all([
        knex.schema.createTable('payment_category', function(table) {
            table
                .increments('cid')
                .unsigned()
                .primary();
            table
                .integer('category_id')
                .unsigned()
                .notNullable();
            table.string('category_description').notNullable();
            table.integer('sub_category_id').unsigned();
            table.string('sub_category_description');
            table.timestamp('created_at').defaultTo(knex.fn.now());
            table.timestamp('closed_at');
        })
    ]);
};
*/
