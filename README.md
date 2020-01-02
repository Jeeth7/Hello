# Hello
Go experiments
This repository is created in conjuction to the youtube content Lear go lang in 7 hours.

struct member{
id
name 
gender}

Branch{
id
MotherId
SpouseId
}

ADD_CHILD(mother name ,child name, gender){
motherid = find mother (mother name)
childid = Create a new member with the childname and Gender
Create a branch object with motherid,childid as id
}

GET_RELATIONSHIP(name, relationship)
{
Identify the name from the memberlist

Paternal-Uncle:
	motherid = get mother id
	father id = get spouse id from mother id
	grand mother id = the mother id of father id
	Get all the child of grand mother id who are male
	
Maternal-Uncle 
	motherid = get mother id
	grand mother id = the mother id of mother id
	Get all the child of grand mother id who are male

Paternal-Aunt 
Maternal-Aunt 
Sister-In-Law 
Brother-In-Law 
Son 
Daughter 
Siblings
}

	
