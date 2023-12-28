-- init database

CREATE TABLE todos (
  	id INT GENERATED ALWAYS AS IDENTITY,
	name VARCHAR ( 50 )
);

CREATE TABLE todos_ref (
	id INT,
	id_ref INT
);

-- init data

INSERT INTO todos(name)
VALUES('a'), ('b'), ('c');

INSERT INTO todos_ref(id, id_ref)
VALUES(1, 2), (1, 3) ,(2,3);

-- topological_graph

WITH
	RECURSIVE recursive_depth_calculator as (
		SELECT 
			id_ref, 
			0 as depth
 		FROM todos_ref WHERE id = 1
 		
		UNION ALL
		
		SELECT
			e.id_ref,
			depth + 1
		FROM
			todos_ref e
		INNER JOIN recursive_depth_calculator s ON s.id_ref = e.id
	),
	result as (
		SELECT * FROM recursive_depth_calculator 
	)
SELECT id_ref, MAX(depth) FROM result GROUP by id_ref ORDER by max