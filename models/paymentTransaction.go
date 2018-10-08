package models

/**
  payment_transaction          // 결제 완료 테이블(charge). 파라미터 비교시 payment_try 내용과 다를경우 hacking으로 간주
     PxID                         // payment_try의 pid
     transaction_id              // pg사로 부터 넘어오는 unique id로 pg사 이용해서 추적이 가능해야 한다.
     user_id
     item_id
     pg_id
     currency                    // default: 'USD'.
     price
     amount                      //cyber coin amount
     transaction_at              // 결제 완료일
     amount_after_used           // 사용 후 남은 amount (insert시 충전되는 amount 와 동일하게...deduct 뙬때 마이너스)
     is_canceled                 // default: 0(false). 향후 cancel 발생을 대비. 향후 cancel 이력 관련 테이블 필요
     canceled_at                //
*/
