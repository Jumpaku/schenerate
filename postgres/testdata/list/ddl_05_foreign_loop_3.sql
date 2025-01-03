-- classDiagram
--     F_1 <|-- F_2
--     F_2 <|-- F_3
--     F_3 <|-- F_1
CREATE TABLE "F_1" (
    "PK_11" integer NOT NULL,
    "PK_12" integer NOT NULL,
    PRIMARY KEY ("PK_11", "PK_12")
);

CREATE TABLE "F_2" (
    "PK_21" integer NOT NULL,
    "PK_22" integer NOT NULL,
    CONSTRAINT "FK_F_2_1" FOREIGN KEY ("PK_21", "PK_22") REFERENCES "F_1" ("PK_11", "PK_12"),
    PRIMARY KEY ("PK_21", "PK_22")
);

CREATE TABLE "F_3" (
    "PK_31" integer NOT NULL,
    "PK_32" integer NOT NULL,
    CONSTRAINT "FK_F_3_2" FOREIGN KEY ("PK_31", "PK_32") REFERENCES "F_2" ("PK_21", "PK_22"),
    PRIMARY KEY ("PK_31", "PK_32")
);

ALTER TABLE "F_1" ADD CONSTRAINT "FK_F_1_3" FOREIGN KEY ("PK_11", "PK_12") REFERENCES "F_3" ("PK_31", "PK_32");