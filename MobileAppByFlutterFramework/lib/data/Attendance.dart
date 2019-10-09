class Attendance{
  final int id;
  final String idStudent;
  final String nameStudent;
  final List<AuthenStudents> authenStudent;

 



  // Attendance({int id,String idStudent,String nameStudent,List<AuthenStudents> authenStudent}){
  //     this.id = id;
  //     this.idStudent = idStudent;
  //     this.nameStudent = nameStudent;
  //     this.authenStudent = authenStudent;
  // }
  Attendance({this.id,this.idStudent,this.nameStudent,this.authenStudent});

  factory Attendance.fromJson(Map<String, dynamic> json){
    var list = json['authenStudents'] as List;
    print(list.runtimeType);
    List<AuthenStudents> subjectsList = list.map((i) => AuthenStudents.fromJson(i)).toList();
    return Attendance( 
      id: json['id'],
      idStudent: json['idStudent'],
      nameStudent: json['nameStudent'],
      authenStudent: subjectsList
    );
  }

  // Map toJson(){
  //   return{
  //     'id' : id,
  //     'idStudent' : idStudent,
  //     'nameStudent' : nameStudent,

  //   };
  // }

  

}



class AuthenStudents{
 final String data;
 final bool stateAuthen;

AuthenStudents({this.data,this.stateAuthen});
  // AuthenStudents({String data,String stateAuthen}){
  //   this.data = data;
  //   this.stateAuthen = stateAuthen;
  // }

  factory AuthenStudents.fromJson(Map<String, dynamic>json){
    return AuthenStudents( 
      data: json['date'],
      stateAuthen: json['stateAuthen']
    );
  }

  // Map toJson(){
  //   return{
  //     'data':data,
  //     'stateAuthen':stateAuthen

  //   };
  // }

}