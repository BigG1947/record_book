package db

// PEOPLE
const TablePeople = "people"
const FieldFIO = "fio"
const FieldBirthaday = "birthday"
const FieldGender = "gender"
const FieldComment = "comment"
const FieldPassword = "password"
const FieldPhoneNumber = "phone_number"
const FieldImg = "img"
const FieldEmail = "email"
const FieldIdStatus = "id_status"
const FieldHaveAccess = "have_access"

// STUDENT
const TableStudent = "student"
const FieldDateAdmission = "date_admission"
const FieldIsFullTime = "is_full_time"
const FieldIsCut = "is_cut"
const FieldIdGroup = "id_group"
const FieldSemester = "semester"

// EMPLOYEE
const TableEmployee = "employee"
const FieldDateInvite = "date_invite"
const FieldIdRank = "id_rank"
const FieldIdCathedra = "id_cathedra"

// GROUPS
const TableGroups = "groups"
const FieldTableName = "name"
const FieldTableIdEmployee = "id_employee"
const FieldTableIdDirection = "id_direction"

// RANK
const TableRank = "ranks"
const FieldRankName = "name"

// DIRECTION
const TableDirection = "direction"
const FieldDirectionName = "name"
const FieldDirectionIdCathedra = "id_cathedra"

// CATHEDRA
const TableCathedra = "cathedra"
const FieldCathedraName = "name"
const FieldCathedraIdFaculty = "id_faculty"

// FACULTY
const TableFaculty = "faculty"
const FieldFacultyName = "name"
const FieldFacultyIdInstitute = "id_institute"

// INSTITUTE
const TableInstitute = "institute"
const FieldInstituteName = "name"

// LOADS
const TableLoads = "loads"
const FieldLoadsId = "id"
const FieldLoadsIdDiscipline = "id_discipline"
const FieldLoadsIdEmployee = "id_employee"
const FieldLoadsIdGroup = "id_group"
const FieldLoadsIdAssistant = "id_assistant"
const FieldLoadsIdSemester = "semester"
