-- Statements that the read-only guard must accept.
SELECT COUNT(*) FROM PCFILIAL;
select codigo, razaosocial from pcfilial where rownum < 10;
WITH t AS (SELECT 1 AS n FROM DUAL) SELECT n FROM t;
/* leading comment */ SELECT 'a;b' AS s FROM DUAL;
-- comment with UPDATE inside
SELECT 'update pcmov set x = 1' AS txt FROM DUAL;
