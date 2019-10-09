
class Subjects{
 String sName;
 int tspassword;

Subjects({String sName,int tspassword}){
  this.sName = sName;
  this.tspassword = tspassword;
}
 factory Subjects.fromJson(Map<String, dynamic> json){
   return Subjects(
    sName: json['TSName'],
    tspassword: json['TSpassword'],
   );
  }

Map tojson(){
  return {
      'TSName': sName, 
      'TSpassword':tspassword,
  };
}

}
