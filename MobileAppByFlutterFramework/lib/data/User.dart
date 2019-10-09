

class User {
  String id;
  String firstname;
  String lastname;
  String idstudent;
  String phone;
  String email;
  String username;
  String password;
  // List<Subjects> subjects;

  User({String id,String firstname, String lastname,String idstudent,String phone,String email,String username, String password /*List<Subjects> subjects*/}) {
    this.id = id;
    this.firstname = firstname;
    this.lastname = lastname;
    this.idstudent = idstudent;
    this.phone = phone;
    this.email = email;
    this.username = username;
    this.password = password;

    // this.subjects = subjects;
  }
 
  factory User.fromJson(Map<String, dynamic> json){
    // var list = json['subjects'] as List;
    // print(list.runtimeType);
    // List<Subjects> subjectsList = list.map((i) => Subjects.fromJson(i)).toList();
    return User(
      id: json['student_id'],
      idstudent: json['SID'],
      firstname: json['SfirstName'],
      lastname:json['SlastName'],
      username :json['Susername'],
      password : json ['Spassword'],
      phone:json['Sphone'],
      email:json['Semail']
      
      // subjects: subjectsList
    );
  }
         

  Map toJson() {
    return {
      'student_id': id,
      'SID': idstudent, 
      'SfirstName': firstname, 
      'SlastName': lastname, 
      'Susername':username,
      'Spassword':password,
      'Sphone': phone, 
      'Semail': email
      
    };
  }
}
