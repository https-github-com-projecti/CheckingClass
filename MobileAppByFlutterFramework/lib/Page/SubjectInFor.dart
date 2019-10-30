import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:project1/data/API.dart';
import 'package:project1/data/Attendance.dart';
import 'package:project1/main.dart';
  String date = " ";
  List<Map<String, IconData>> icon = [
    {
      'false': Icons.close ,
      'true': Icons.done,
    }, 
  ];
  Attendance att;
  

class SubjectInFor extends StatefulWidget {
  static const String routeName = "/subjectinfor";
  final  attendance;
  final idstudent;
  final tspassword;
 

  const SubjectInFor({Key key, this.attendance,this.idstudent,this.tspassword}) : super(key: key);
  @override
  Episode5State createState() {   
     att = attendance;
    return new Episode5State();
  }
}

 
class Episode5State extends State<SubjectInFor> {
  var status;
  final GlobalKey<RefreshIndicatorState> _refreshIndicatorKey = new GlobalKey<RefreshIndicatorState>();
    @override
  void initState() {
    super.initState();
   
    WidgetsBinding.instance
        .addPostFrameCallback((_) => _refreshIndicatorKey.currentState.show());
    
  }

 
  Widget bodyData() => DataTable(
     columns: <DataColumn>[
              DataColumn(
                
                label: Text("Date and Time",
                 style: TextStyle( 
                   color: Colors.black,
                   fontSize: 16.0,
                 )

                )
                
              ),
              DataColumn(
                label: Text("Attendance",
                  style: TextStyle( 
                    color: Colors.black,
                    fontSize: 16.0,
                    
                  )
                ),
              ),
      ],
    rows: [
     
     
      for (var i = 0 ;i < att.authenStudent.length;i++) 
         DataRow(
                
                cells: <DataCell>[
                  DataCell(Text(att.authenStudent[i].date[0]+att.authenStudent[i].date[1]
                  +att.authenStudent[i].date[2]+att.authenStudent[i].date[3]
                  +att.authenStudent[i].date[4]+att.authenStudent[i].date[5]+
                  att.authenStudent[i].date[6]+att.authenStudent[i].date[7]+
                  att.authenStudent[i].date[8]+att.authenStudent[i].date[9]
                  +att.authenStudent[i].date[10]+att.authenStudent[i].date[11])),
                  DataCell(Icon(icon[0][att.authenStudent[i].stateAuthen.toString()]))
                
                ],
          ),
          
    ],
           
  );

  Future<void> getDataSubjectData() async{
    API.getAttendance(widget.idstudent,widget.tspassword.toString()).then((response) {   
      print(response.body);
      print(json.decode(response.body));
        Map list = json.decode(response.body);
        atten = new Attendance.fromJson(list);
        final int statusCode = response.statusCode;
        if (statusCode < 200 || statusCode > 400 || json == null) {
          throw new Exception("Error while fetching data");
        }
      });
      setState(() {
         att = atten;
        });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Attendance Table"),
        backgroundColor: Colors.amber[numColor],
      ),
      body: RefreshIndicator(
        onRefresh: getDataSubjectData,
        child: SafeArea(
          
          child: Container(
            child: bodyData(),
          ),
        ),
      ),
    );
  }
}
