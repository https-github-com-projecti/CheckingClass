import 'dart:convert';
import 'dart:io';
import 'package:camera/camera.dart';
import 'package:flutter/material.dart';
import 'dart:async';
import 'package:flutter_barcode_scanner/flutter_barcode_scanner.dart';
import 'package:flutter_image_compress/flutter_image_compress.dart';
import 'package:image_picker/image_picker.dart';
import 'package:project1/data/API.dart';
import 'package:project1/data/Attendance.dart';
import 'package:project1/data/DataCheck.dart';
import 'package:project1/data/Selfie.dart';
import 'package:project1/data/Subjects.dart';
import 'package:project1/data/User.dart';
import 'package:project1/main.dart';
import 'SubjectInFor.dart';




List<CameraDescription> cameras;
File file;

class Subject extends StatefulWidget{
   static const String routeName = "/subject";
  final User user ;
  final int tspassword;
  Subject({Key key, @required this.user,this.tspassword }) : super(key: key);
  
  @override
  _SubjectState createState() => _SubjectState();
  
}

class _SubjectState extends State<Subject> {
 

@override
  Widget build(BuildContext context){
    return Scaffold( 
      appBar: AppBar( 
        elevation: 0,
        backgroundColor: Colors.amber,
        iconTheme: new IconThemeData(color: Colors.black),
      ),
      backgroundColor: Colors.amber,
      body: Center( 

        child: Container(
          margin: const EdgeInsets.all(50.0),
          child: Column( 
            children: <Widget>[ 
              GestureDetector(
                  onTap: () {
                    print("Card 1");
                    // print(widget.user.idstudent);
                    print(attendance[0].authenStudent);

                    // getAttendance();
                    // print(atten.nameStudent);
                    Navigator.push(context,MaterialPageRoute(builder: (context) => SubjectInFor(attendance: attendance[0],idstudent: widget.user.idstudent,tspassword: widget.tspassword,)));
                    // Navigator.of(context).pushNamed("/" + camera);         
                    // Navigator.of(context).pushNamed("/"+ subjectInFor);           
                  },
                  child: Card(                    
                    elevation: 5.0,
                    child: Container(
                      width: 250,
                      height: 200,
                      child: Column(  
                        children: <Widget>[ 
                          Padding(
                            padding: const EdgeInsets.only(top: 70),
                            child: Text("AttendanceTable",
                              style: TextStyle(
                                fontFamily: "Poppins-Bold",
                                fontSize: 25,
                                fontWeight: FontWeight.bold,
                              )
                             ),
                          )
                        ],
                      ),
                    ),

                  ),
              ),
              GestureDetector(
                  onTap: () async {
                    print("Card 2");
                     _incrementCounter();
                    
                     
        
                  },
                  child: Card(
                    
                    elevation: 5.0,
                    child: Container(
                      width: 250,
                      height: 200,
                      child: Column(  
                        children: <Widget>[ 
                          Padding(
                            padding: const EdgeInsets.all(60.0),
                            child: Text("scan ",
                              style: TextStyle(
                                fontFamily: "Poppins-Bold",
                                fontSize: 50,
                                fontWeight: FontWeight.bold,
                              )
                             ),
                          )
                        ],
                      ),
                    ),
                  ),
              )   
            ],
          ),
        ),
      )
    );
  }
@override
  void initState() {
    super.initState();
    getAttendance();
  }
  String _counter,_value = "";
  String x ="AAAAAAAA";

  Future _incrementCounter() async{
    _counter= await FlutterBarcodeScanner.scanBarcode("#11f73f", "Cancel", true, ScanMode.DEFAULT);
    setState(() {_value = _counter;});
    if(_value != "-1"){
      
      DataCheck newSend = new DataCheck(dataCheck: _value,tspassword: widget.tspassword,idstudent: widget.user.idstudent);
      DataCheck news = await send(dateCheck,body: newSend.toMap());  
      if (news != null){
        print("--------------NULL-------------");
      }
      print(scanState.toString());
      print("_______________");
      if(scanState.toString() == "Success"){
        _choose();
      }else if(scanState.toString() == "Error"){
        showDialog(             
          context: context,
          builder: (context) => new AlertDialog(
              shape: new RoundedRectangleBorder(borderRadius: new BorderRadius.circular(10.0)),
              title: new Text("Can't Selfie.",
                textAlign: TextAlign.center,
              ),
              content: Container(
                  height: 90.0, 
                  child: Column(  
                    children: <Widget>[ 
                      Icon(
                        Icons.not_interested,
                        color: Colors.red,
                        size: 70.0,
                      
                      )
                    ],
                  ),
                ),                   
              )
            );
      }
      
      
    } 
  }

   void _choose() async {
  //  file = await ImagePicker.pickImage(source: ImageSource.camera);
      file = await ImagePicker.pickImage(source: ImageSource.camera,imageQuality: 25);
        _upload();
   }
  void getAttendance() async {
   
    API.getAttendance(widget.user.idstudent,widget.tspassword.toString()).then((response) {   
      print(response.body); 
      print("------------------------------------------------------------------------");
      print(json.decode(response.body));
        Iterable<dynamic> list = json.decode(response.body); 
        attendance = list.map((model) => Attendance.fromJson(model)).toList();
        // Iterable<dynamic> list = json.decode(response.body);
        // Iterable<AuthenStudents> list = json.decode(response.body);
        
         
        // Map list = json.decode(response.body);
        // atten = new Attendance.fromJson(list);

        // final int statusCode = response.statusCode;
        // if (statusCode < 200 || statusCode > 400 || json == null) {
        //   throw new Exception("Error while fetching data");
        // }
      });
  }
      
    

      



  Future<void> _upload() async {
   if (file == null) return;
  //  testCompressAndGetFile(file, targetPath)
   String base64Image = base64Encode(file.readAsBytesSync());
    print("<<<<<<<<<<<<|||>>>>>>>>>");
    print(base64Image);
   Selfie newSend = new Selfie(studentId: widget.user.idstudent,imageSelfie: base64Image,dataCheck: _value,firstName: widget.user.firstname,lastName: widget.user.lastname);
   Selfie news = await selfieQR(selfieqr,body: newSend.toMap());
    if (news != null){
      print("--------------NULL-------------");
    }
    if(selfieState == "Success"){
      showDialog(             
          context: context,
          builder: (context) => new AlertDialog(
              shape: new RoundedRectangleBorder(borderRadius: new BorderRadius.circular(10.0)),
              title: new Text("  Success  ",
                textAlign: TextAlign.center,
              ),
              content: Container(
                  height: 90.0, 
                  child: Column(  
                    children: <Widget>[ 
                      Icon(
                        Icons.done_outline,
                        color: Colors.green,
                        size: 70.0,
                      
                      )
                    ],
                  ),
                ),                   
              )
            );
    }
 }
      
  Future<File> testCompressAndGetFile(File file, String targetPath) async {
    var result = await FlutterImageCompress.compressAndGetFile(
        file.absolute.path, targetPath,
        quality: 88,
        rotate: 180,
      );
    print(file.lengthSync());
    print(result.lengthSync());
    return result;
  }
         
      
  
}