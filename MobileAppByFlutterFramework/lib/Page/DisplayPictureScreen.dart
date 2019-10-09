// import 'dart:convert';
// import 'dart:io';
// import 'package:flutter/material.dart';
// import 'package:project1/Page/SubjectInFor.dart';
// import 'package:project1/data/Selfie.dart';
// import 'package:project1/main.dart';

// import 'Subject.dart';




// class DisplayPictureScreen extends StatefulWidget {
//   final String imagePath;
//   final String studentId;
//   final String firstname;
//   final String lastname;
//   final String image;
//   final String qr;
//   const DisplayPictureScreen({Key key, this.imagePath,this.studentId,this.image,this.qr,this.firstname,this.lastname}) : super(key: key);

//   @override
//   _DisplayPictureScreenState createState() => _DisplayPictureScreenState();
// }

// class _DisplayPictureScreenState extends State<DisplayPictureScreen> {

  

//   @override
//   Widget build(BuildContext context) {    

//     return Scaffold(
//       appBar: AppBar(title: Text('Display the Picture')),
//       body: FutureBuilder<void>(
//         // future: _initializeControllerFuture,
//         builder: (context, snapshot) {
//           // if (snapshot.connectionState == ConnectionState.done) {   // ตัวอย่าง รอโหลด !!!
//             return Image.file(File(widget.imagePath),height: 1080,width: 1000,);
//           // } else {
//           //   return Center(child: CircularProgressIndicator());
//           // }
//         },
//       ),
//      floatingActionButton: FloatingActionButton(
//         child: Icon(Icons.send),
//         onPressed: () async {
//           var bytes = utf8.encode(widget.image);
//           String image64 = base64Encode(bytes);
//           Selfie newSend = new Selfie(studentId: widget.studentId,imageSelfie: image64,dataCheck: widget.qr,firstName: widget.firstname,lastName: widget.lastname);
//           Selfie news = await selfieQR(selfieqr,body: newSend.toMap());
//           Navigator.of(context).pushNamed("/" + subject);
//             // Navigator.push(context,MaterialPageRoute(builder: (context) => Subject()));
          
         
//         AlertDialog();
//         }
//       ),

//     );
        
  

//   }  
// }
