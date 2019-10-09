import 'package:flutter/material.dart';
import 'package:project1/data/Attendance.dart';
import 'package:project1/main.dart';

class SubjectInFor extends StatefulWidget {
  static const String routeName = "/subjectinfor";
  final  attendance;

  const SubjectInFor({Key key, this.attendance}) : super(key: key);
  @override
  Episode5State createState() {
    return new Episode5State();
  }
}

class Episode5State extends State<SubjectInFor> {
  Widget bodyData() => DataTable(
     columns: [
              DataColumn(
                label: Text("Date and Time"),
              ),
              DataColumn(
                label: Text("Attendance"),
              ),
            ],
    rows: [
      for (var i = 0 ;i < widget.attendance.authenStudent.length;i++) 
         DataRow(
                
                cells: [
                  DataCell(Text(widget.attendance.authenStudent[i].data.toString())),
                  DataCell(Text(widget.attendance.authenStudent[i].stateAuthen.toString())),
                
                ],
              ),
            ],
           
      );

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Attendance Table"),
        backgroundColor: Colors.amber[numColor],
      ),
      body: Container(
        child: bodyData(),
      ),
    );
  }
}
