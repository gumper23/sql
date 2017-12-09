create table if not exists ints
(
        si smallint
        , i int
        , bi bigint
);

create or replace function get_random_number(in_min bigint, in_max bigint) returns bigint as $$ 
begin 
    return floor(random() * (in_max - in_min + 1)) + in_min; 
end; 
$$ language plpgsql;

create or replace function populate_ints(in_num int) returns void  as $$
declare
    i int := 0;
begin
    for i in 1..in_num loop
        insert into ints(si, i, bi)
        values (
            get_random_number(-32768, 32767)
            , get_random_number(-2147483648, 2147483647)
            , get_random_number(-2305843009213693951, 2305843009213693951)
        );
    end loop;
end;
$$ language plpgsql;

select populate_ints(100);

