create database if not exists qm;
use qm;

create table if not exists qm.ints
(
	ti			tinyint 				not null 	default '0'
	, tiu		tinyint unsigned 		not null 	default '0'
	, si		smallint 				not null	default '0'
	, siu		smallint unsigned 		not null 	default '0'
	, mi 		mediumint 				not null 	default '0'
	, miu 		mediumint unsigned 		not null 	default '0'
	, i 		int 					not null 	default '0'
	, iu 		int unsigned 			not null 	default '0'
	, bi 		bigint 					not null 	default '0'
	, biu 		bigint unsigned 		not null 	default '0'
);

## MySQL function to generate random numbers
DELIMITER $$
DROP FUNCTION IF EXISTS qm.get_random_number $$
CREATE FUNCTION qm.get_random_number(in_min bigint, in_max bigint) RETURNS BIGINT DETERMINISTIC
    RETURN FLOOR(RAND() * (in_max - in_min + 1)) + in_min
$$
DELIMITER ;

DELIMITER $$
DROP PROCEDURE IF EXISTS qm.populate_ints $$
CREATE PROCEDURE qm.populate_ints(in_numrows int)
BEGIN
	DECLARE i int;
	SET i := 0;

	WHILE (i < in_numrows) DO
		insert into ints(ti, tiu, si, siu, mi, miu, i, iu, bi, biu)
		values (
			get_random_number(-128, 127) 									-- ti
			, get_random_number(0, 255) 									-- tiu
			, get_random_number(-32768, 32767)								-- si
			, get_random_number(0, 65535)									-- siu
			, get_random_number(-8388608, 8388607) 							-- mi
			, get_random_number(0, 16777215) 								-- miu
			, get_random_number(-2147483648, 2147483647) 					-- i
			, get_random_number(0, 4294967295) 								-- iu
			, get_random_number(-2305843009213693951, 2305843009213693951) 	-- bi
			, get_random_number(0, 2305843009213693951) 					-- biu
		);
		SET i := i + 1;		
	END WHILE;
END $$
DELIMITER ;

insert into ints(ti, tiu, si, siu, mi, miu, i, iu, bi, biu)
values 
(
	-128
	, 0
	, -32768
	, 0
	, -8388608
	, 0
	, -2147483648
	, 0
	, -9223372036854775808
	, 0
);

insert into ints(ti, tiu, si, siu, mi, miu, i, iu, bi, biu)
values 
(
	127
	, 255
	, 32767
	, 65535
	, 8388607
	, 16777215
	, 2147483647
	, 4294967295
	, 9223372036854775807
	, 18446744073709551615
);

call populate_ints(100);

create table if not exists dates
(
    d           date
    , dt        datetime
    , ts        timestamp
    , tm        time
    , yr        year
);

DELIMITER $$
DROP PROCEDURE IF EXISTS qm.populate_dates $$
CREATE PROCEDURE qm.populate_dates(in_numrows int)
BEGIN
	DECLARE i int;
	DECLARE dt datetime;
	SET i := 0;

	WHILE (i < in_numrows) DO
		set dt := from_unixtime(get_random_number(0, 2147483647));
		insert into dates(d, dt, ts, tm, yr)
		values (
			date(dt)				-- d
			, dt  					-- dt
			, dt					-- ts
			, time(dt)  			-- tm
			, year(dt)  			-- yr
		);
		SET i := i + 1;		
	END WHILE;
END $$
DELIMITER ;

call populate_dates(100);




/*
create table if not exists numbers
(
	id 			serial 
	, b			bit 					not null	default 0
	, t 		tinyint 				not null 	default '0'
	, tu		tinyint unsigned		not null 	default '0'
	, bl		boolean 				not null 	default '0'
	, si		smallint 				not null	default '0'
	, siu 		smallint unsigned 		not null 	default '0'
	, i 		int 					not null 	default '0'
	, iu 		int unsigned 			not null 	default '0'
	, d 		decimal(8,2) 			not null 	default '0.0'
	, du 		decimal(8,2) unsigned 	not null 	default '0.0'
	, f 		float					not null 	default '0.0'
	, fu 		float unsigned 			not null 	default '0.0'
	, db		double					not null 	default '0.0'
	, dbu		double unsigned 		not null 	default '0.0'
);

insert into numbers(b) values (true);
insert into numbers(b) values (false);

create table if not exists dates
(
	d 			date
	, dt		datetime
	, ts		timestamp
	, tm		time
	, yr		year
);

insert into dates(d, dt, ts, tm, yr) values(curdate(), now(), current_timestamp(), curtime(), year(now()));
*/
