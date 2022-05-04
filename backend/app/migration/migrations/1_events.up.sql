create table if not exists events (
    'id' integer auto_increment primary key,
    'title' varchar(40),
    'start' datetime,
    'end' datetime
)