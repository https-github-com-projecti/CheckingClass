// import 'dart:async';
// import 'dart:convert';
// import 'dart:io';
// import 'package:camera/camera.dart';
// import 'package:flutter/material.dart';
// import 'package:flutter/services.dart';
// import 'package:image_picker/image_picker.dart';
// import 'package:path/path.dart' show join;
// import 'package:path_provider/path_provider.dart';
// import 'package:project1/data/User.dart';
// import 'package:project1/meterial/fab_bottom_app_bar.dart';
// import 'DisplayPictureScreen.dart';

// int i=1;
// File file;
// var whatTheFile;
// class TakePictureScreen extends StatefulWidget {
//   final List<CameraDescription> camera;
//   final User user;
//   final String qr;
//   const TakePictureScreen({Key key,@required this.camera,this.user,this.qr}) : super(key: key);

//   @override
//   TakePictureScreenState createState() => TakePictureScreenState();
// }

// class TakePictureScreenState extends State<TakePictureScreen> {
//   CameraController _controller;
//   Future<void> _initializeControllerFuture;
//   String _lastSelected = 'TAB: 0';
  
//   @override
//   void initState() {
//     super.initState();
      
//     _controller = CameraController(
//       widget.camera[i],
//       ResolutionPreset.high,
//     );

//     // Next, initialize the controller. This returns a Future.
//     _initializeControllerFuture = _controller.initialize();
//   }

//   @override
//   void dispose() {
//     // Dispose of the controller when the widget is disposed.
//     _controller.dispose();
//     super.dispose();
//   }

//   @override
//   Widget build(BuildContext context) {
//     SystemChrome.setPreferredOrientations([
//         DeviceOrientation.portraitUp,
//         DeviceOrientation.portraitDown,
//     ]);
    
//     return Scaffold(
//       // appBar: AppBar(title: Text('Take a picture')),
//       body: FutureBuilder<void>(
//         future: _initializeControllerFuture,
//         builder: (context, snapshot) {
//           if (snapshot.connectionState == ConnectionState.done) {   // ตัวอย่าง รอโหลด !!!
//             return CameraPreview(_controller);
//           } else {
//             return Center(child: CircularProgressIndicator());
//           }
//         },
//       ),
//       // bottomNavigationBar: BottomAppBar(
//       //   shape: const CircularNotchedRectangle(),
//       //   child: Container(height: 50.0,),
//       // ),
//       bottomNavigationBar: FABBottomAppBar(
//         centerItemText: '',
//         color: Colors.grey,
//         selectedColor: Colors.red,
//         notchedShape: CircularNotchedRectangle(),
//         onTabSelected: _selectedTab,
//         items: [
//           FABBottomAppBarItem(iconData: Icons.arrow_back_ios, text: 'Back'),
//           // FABBottomAppBarItem(iconData: Icons.layers, text: 'Is'),
//           FABBottomAppBarItem(iconData: Icons.switch_camera, text: 'Switch Camera'),
//           // FABBottomAppBarItem(iconData: Icons.info, text: 'Bar'),
//         ],
//       ),
//       floatingActionButton: FloatingActionButton(
//         child: Icon(Icons.camera_alt),
//         onPressed: () async {
//           try {
//             await _initializeControllerFuture;
//             final path = join(
//               (await getTemporaryDirectory()).path,
//               '${DateTime.now()}.png',
//             );
            


//             String imageFile = '${DateTime.now()}.png';    //No Directory
//             await _controller.takePicture(path);

//             whatTheFile = Image.file(File(path));   ////
            
//             _choose();
//             _upload();
         
             
//             Navigator.push(
//               context,
//               MaterialPageRoute(builder: (context) => DisplayPictureScreen(imagePath: path,image: imageFile,studentId: widget.user.idstudent,qr: widget.qr,firstname: widget.user.firstname,lastname: widget.user.lastname,),),
//             );
//             // print(imageFile+"***************************************************");
//             print(path);
//           } catch (e) {print(e);}
//         },
//       ),
//       floatingActionButtonLocation: FloatingActionButtonLocation.centerDocked,
//     );
//   }



   

//   void _selectedTab(int index) {

//     setState(() {
//       _lastSelected = 'TAB: $index';
//       print(_lastSelected);
//     });
//   }

//   void _choose() async {
//   //  file = await ImagePicker.pickImage(source: ImageSource.camera);
//   file = await ImagePicker.pickImage(source: ImageSource.camera);
//  }

//   void _upload() {
//    if (file == null) return;
//    String base64Image = base64Encode(file.readAsBytesSync());
//    String fileName = file.path.split("/").last;
//     print("<<<<<<<<<<<<>>>>>>>>>");
//     print("<<<<<<<<<<<<>>>>>>>>>");
//     print("<<<<<<<<<<<<>>>>>>>>>");
//    print(base64Image);

   
//  }
// }

