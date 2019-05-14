package db

const getAllPeopleScript = `SELECT people.id,
       people.fio,
       people.birthday,
       people.gender,
       people.img,
       people.comment,
       people.password,
       people.phone_number,
       people.email,
       people.have_access,
       student.id_people,
       student.date_admission,
       student.is_full_time,
       student.is_cut,
       SG.id,
       SG.id_employee,
       SG.name,
       SG.id_direction,
       student.semester,
       employee.id_people,
       employee.date_invite,
       ranks.id,
       ranks.name,
       cathedra.id,
       cathedra.name,
       cathedra.id_faculty,
       accession.edit_access,
       accession.set_absence,
       accession.get_absence,
       accession.set_mark,
       accession.set_event,
       accession.get_sensitive,
       accession.set_sensitive,
       accession.get_ylist,
       accession.manage_academ,
       sensitive_data.passport_code,
       sensitive_data.rntrs,
       sensitive_data.reg_address,
       sensitive_data.military_id,
       status.id,
       status.name
FROM people
       LEFT OUTER JOIN student ON people.id = student.id_people
       LEFT OUTER JOIN record_book.groups SG ON student.id_group = SG.id
       LEFT OUTER JOIN employee ON employee.id_people = people.id
       LEFT OUTER JOIN ranks ON employee.id_rank = ranks.id
       LEFT OUTER JOIN cathedra ON employee.id_cathedra = cathedra.id
       INNER JOIN accession ON people.id = accession.id_people
       INNER JOIN sensitive_data ON people.id = sensitive_data.id_people
       INNER JOIN status ON people.id_status = status.id;`

const getPeopleByIdScript = `SELECT people.id,
       people.fio,
       people.birthday,
       people.gender,
       people.img,
       people.comment,
       people.password,
       people.phone_number,
       people.email,
       people.have_access,
       student.id_people,
       student.date_admission,
       student.is_full_time,
       student.is_cut,
       SG.id,
       SG.id_employee,
       SG.name,
       SG.id_direction,
       student.semester,
       employee.id_people,
       employee.date_invite,
       ranks.id,
       ranks.name,
       cathedra.id,
       cathedra.name,
       cathedra.id_faculty,
       accession.edit_access,
       accession.set_absence,
       accession.get_absence,
       accession.set_mark,
       accession.set_event,
       accession.get_sensitive,
       accession.set_sensitive,
       accession.get_ylist,
       accession.manage_academ,
       sensitive_data.passport_code,
       sensitive_data.rntrs,
       sensitive_data.reg_address,
       sensitive_data.military_id,
       status.id,
       status.name
FROM people
       LEFT OUTER JOIN student ON people.id = student.id_people
       LEFT OUTER JOIN record_book.groups SG ON student.id_group = SG.id
       LEFT OUTER JOIN employee ON employee.id_people = people.id
       LEFT OUTER JOIN ranks ON employee.id_rank = ranks.id
       LEFT OUTER JOIN cathedra ON employee.id_cathedra = cathedra.id
       INNER JOIN accession ON people.id = accession.id_people
       INNER JOIN sensitive_data ON people.id = sensitive_data.id_people
       INNER JOIN status ON people.id_status = status.id
WHERE people.id = (?);`

const getPeopleByEmailScript = `SELECT people.id,
       people.fio,
       people.birthday,
       people.gender,
       people.img,
       people.comment,
       people.password,
       people.phone_number,
       people.email,
       people.have_access,
       student.id_people,
       student.date_admission,
       student.is_full_time,
       student.is_cut,
       SG.id,
       SG.id_employee,
       SG.name,
       SG.id_direction,
       student.semester,
       employee.id_people,
       employee.date_invite,
       ranks.id,
       ranks.name,
       cathedra.id,
       cathedra.name,
       cathedra.id_faculty,
       accession.edit_access,
       accession.set_absence,
       accession.get_absence,
       accession.set_mark,
       accession.set_event,
       accession.get_sensitive,
       accession.set_sensitive,
       accession.get_ylist,
       accession.manage_academ,
       sensitive_data.passport_code,
       sensitive_data.rntrs,
       sensitive_data.reg_address,
       sensitive_data.military_id,
       status.id,
       status.name
FROM people
       LEFT OUTER JOIN student ON people.id = student.id_people
       LEFT OUTER JOIN record_book.groups SG ON student.id_group = SG.id
       LEFT OUTER JOIN employee ON employee.id_people = people.id
       LEFT OUTER JOIN ranks ON employee.id_rank = ranks.id
       LEFT OUTER JOIN cathedra ON employee.id_cathedra = cathedra.id
       INNER JOIN accession ON people.id = accession.id_people
       INNER JOIN sensitive_data ON people.id = sensitive_data.id_people
       INNER JOIN status ON people.id_status = status.id
WHERE people.email = (?);`

const getAllStudentsScript = `SELECT people.id, 
    people.fio,   
    people.birthday,
    people.gender,
    people.img,
    people.comment,
    people.password,
    people.phone_number, 
    people.email, 
    people.have_access,
    student.id_people,
    student.date_admission,
    student.is_full_time,
    student.is_cut,
    groups.id,
    groups.id_employee,
    groups.name,
    groups.id_direction,
    student.semester,
    accession.edit_access,
    accession.set_absence,
    accession.get_absence,
    accession.set_mark,
    accession.set_event,
    accession.get_sensitive,
    accession.set_sensitive,
    accession.get_ylist,
    accession.manage_academ,
    sensitive_data.passport_code,
    sensitive_data.rntrs,
    sensitive_data.reg_address,
    sensitive_data.military_id,
    status.id,
    status.name
FROM people 
	INNER JOIN student ON people.id = student.id_people
	INNER JOIN groups ON student.id_group = groups.id
	INNER JOIN accession ON people.id = accession.id_people
    INNER JOIN sensitive_data ON people.id = sensitive_data.id_people
    INNER JOIN status ON people.id_status = status.id
WHERE people.id IN (SELECT id_people FROM student) ORDER BY status.id ASC, student.id_group DESC, people.fio ASC;`

const getAllEmployeeScript = `SELECT people.id,
       people.fio,
       people.birthday,
       people.gender,
       people.img,
       people.comment,
       people.password,
       people.phone_number,
       people.email,
       people.have_access,
       employee.id_people,
       employee.date_invite,
       ranks.id,
       ranks.name,
       cathedra.id,
       cathedra.name,
       cathedra.id_faculty,
       accession.edit_access,
       accession.set_absence,
       accession.get_absence,
       accession.set_mark,
       accession.set_event,
       accession.get_sensitive,
       accession.set_sensitive,
       accession.get_ylist,
       accession.manage_academ,
       sensitive_data.passport_code,
       sensitive_data.rntrs,
       sensitive_data.reg_address,
       sensitive_data.military_id,
       status.id,
       status.name
FROM people
       INNER JOIN employee ON employee.id_people = people.id
       INNER JOIN cathedra ON employee.id_cathedra = cathedra.id
       INNER JOIN ranks ON employee.id_rank = ranks.id
       INNER JOIN accession ON people.id = accession.id_people
       INNER JOIN sensitive_data ON people.id = sensitive_data.id_people
       INNER JOIN status ON people.id_status = status.id
WHERE people.id IN (SELECT id_people FROM employee) ORDER BY status.id ASC, cathedra.name ASC, ranks.name ASC, people.fio ASC;`

const getStudentFromGroupScript = `SELECT people.id, 
    people.fio,   
    people.birthday,
    people.gender,
    people.img,
    people.comment,
    people.password,
    people.phone_number, 
    people.email, 
    people.have_access,
    student.id_people,
    student.date_admission,
    student.is_full_time,
    student.is_cut,
    groups.id,
    groups.id_employee,
    groups.name,
    groups.id_direction,
    student.semester,
    accession.edit_access,
    accession.set_absence,
    accession.get_absence,
    accession.set_mark,
    accession.set_event,
    accession.get_sensitive,
    accession.set_sensitive,
    accession.get_ylist,
    accession.manage_academ,
    sensitive_data.passport_code,
    sensitive_data.rntrs,
    sensitive_data.reg_address,
    sensitive_data.military_id,
    status.id,
    status.name
FROM people 
	INNER JOIN student ON people.id = student.id_people
	INNER JOIN groups ON student.id_group = groups.id
	INNER JOIN accession ON people.id = accession.id_people
    INNER JOIN sensitive_data ON people.id = sensitive_data.id_people
    INNER JOIN status ON people.id_status = status.id
WHERE people.id IN (SELECT id_people FROM student WHERE id_group = (?)) ORDER BY status.id ASC, student.id_group DESC, people.fio ASC;`

const blockPeopleScript = `UPDATE people SET people.have_access = false WHERE people.id = ?;`

const unblockPeopleScript = `UPDATE people SET people.have_access = true WHERE people.id = ?;`

const updatePeopleStatusScript = `UPDATE people SET people.id_status = ? WHERE people.id = ?;`

const changePeoplePasswordScript = `UPDATE people SET password = ? WHERE people.id = ?`

const changeStudentGroupScript = `UPDATE student SET id_group = ? WHERE student.id_people = ?`

const updateStudentDataScript = `UPDATE student SET date_admission = ?, is_full_time = ?, is_cut = ?, id_group = ?, semester = ? WHERE id_people = ?`

const updateEmployeeDataScript = `UPDATE employee SET date_invite = ?, id_rank = ?, id_cathedra = ? WHERE id_people = ?;`

const updateSensitiveDataScript = `UPDATE sensitive_data SET passport_code = ?, rntrs = ?, reg_address = ?, military_id = ? WHERE sensitive_data.id_people = ?`

const updateAccessionDataScript = `UPDATE accession SET edit_access = ?, set_absence = ?, get_absence = ?, set_mark = ?, set_event = ?, get_sensitive = ?, set_sensitive = ?, get_ylist = ? , manage_academ = ? WHERE id_people = ?;`

const updateGroupInfoScript = `UPDATE groups SET id_employee = ?, name = ?, id_direction = ? WHERE id = ?;`

const deleteGroupScript = `DELETE FROM groups WHERE id = ?;`

const insertRankScript = `INSERT INTO ranks(name) VALUES (?);`

const updateRankScript = `UPDATE ranks SET name = ? WHERE id = ?;`

const deleteRankScript = `DELETE FROM ranks WHERE id = ?;`

const insertDirectionScript = `INSERT INTO direction(name, id_cathedra) VALUES (?, ?);`

const updateDirectionScript = `UPDATE direction SET name = ?, id_cathedra = ? WHERE id = ?;`

const deleteDirectionScript = `DELETE FROM direction WHERE id = ?;`

const insertCathedraScript = `INSERT INTO cathedra(name, id_faculty) VALUES (?, ?);`

const updateCathedraScript = `UPDATE cathedra SET name = ?, id_faculty = ? WHERE id = ?;`

const deleteCathedraScript = `DELETE FROM cathedra WHERE id = ?;`

const insertFacultyScript = `INSERT INTO faculty(name, id_institute) VALUES (?, ?);`

const updateFacultyScript = `UPDATE faculty SET name = ?, id_institute = ? WHERE id = ?;`

const deleteFacultyScript = `DELETE FROM faculty WHERE id = ?;`

const updateInstituteScript = `UPDATE institute SET name = ? WHERE id = ?;`

const insertInstituteScript = `INSERT INTO institute(name) VALUES (?);`

const deleteInstituteScript = `DELETE FROM institute WHERE id = ?;`
