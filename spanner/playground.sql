--sql query unique key information
WITH
    FK_BACKING AS (
        SELECT rc.UNIQUE_CONSTRAINT_NAME AS Name
        FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc
                 JOIN INFORMATION_SCHEMA.REFERENTIAL_CONSTRAINTS rc ON rc.CONSTRAINT_NAME = tc.CONSTRAINT_NAME
                 JOIN INFORMATION_SCHEMA.TABLE_CONSTRAINTS tc2 ON tc2.CONSTRAINT_NAME = rc.UNIQUE_CONSTRAINT_NAME
        WHERE tc.CONSTRAINT_TYPE = 'FOREIGN KEY'
    )
SELECT
    idx.INDEX_NAME AS Name,
    idx.INDEX_NAME IN (SELECT Name FROM FK_BACKING) AS Backing,
    idx.IS_UNIQUE AS IsUnique,
    idx.INDEX_TYPE,
    idx.TABLE_NAME,
    idx.TABLE_SCHEMA,
    ARRAY(
        SELECT AS STRUCT
	    	idxc.COLUMN_NAME AS Name,
            idxc.COLUMN_ORDERING = 'DESC' AS IsDesc
		FROM INFORMATION_SCHEMA.INDEX_COLUMNS idxc
		WHERE idx.INDEX_NAME = idxc.INDEX_NAME AND idx.TABLE_NAME = idxc.TABLE_NAME
		ORDER BY idxc.ORDINAL_POSITION
    ) AS Key,
FROM INFORMATION_SCHEMA.INDEXES idx
WHERE
    idx.TABLE_NAME = 'C_3'
    --idx.TABLE_NAME = 'C_2' --AND idx.INDEX_TYPE != "PRIMARY_KEY"
ORDER BY Name;