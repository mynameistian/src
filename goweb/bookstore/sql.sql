
--创建购物车表

CREATE table Carts(
 id VARCHAR(100) PRIMARY KEY,
 total_count int not null ,
 total_amount numeric(11,2) not null,
 user_id int not null,
 foreign  key(user_id) REFERENCES users(id)
)

--创建购物项表
CREATE TABLE cart_itmes(
	id SERIAL primary key NOT NULL,
	count int not null,
	amount NUMERIC(11,2) NOT NULL,
	book_id int not null,
	cart_id VARCHAR(100) not null,
	foreign  key(book_id) REFERENCES books(id),
	foreign  key(cart_id) REFERENCES carts(id)
)