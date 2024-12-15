--sql query index information
SELECT
    *
FROM pragma_index_list('E') AS pil
    JOIN pragma_index_xinfo(pil.name) AS pii
WHERE pii."name" <> ''
ORDER BY pil."name", pil."seq", pii."seqno"