// import 'package:flutter/material.dart';

// enum DialogState {
//   LOADING,
//   COMPLETED,
//   DISMISSED,
// }
// class MyDialog extends StatelessWidget {
//   final DialogState state;
//   MyDialog({this.state});

//   @override
//   Widget build(BuildContext context) {
//     return state == DialogState.DISMISSED
//         ? Container()
//         : AlertDialog(
//             shape: RoundedRectangleBorder(
//               borderRadius: BorderRadius.all(
//                 Radius.circular(10.0),
//               ),
//             ),
//             content: Container(
//               width: 250.0,
//               height: 100.0,
//               child: state == DialogState.LOADING
//                   ? Row(
//                       mainAxisAlignment: MainAxisAlignment.center,
//                       children: [
//                         CircularProgressIndicator(),
//                         Padding(
//                           padding: const EdgeInsets.only(left: 10.0),
//                           child: Text(
//                             "Loginnig...",
//                             style: TextStyle(
//                               fontFamily: "OpenSans",
//                               color: Color(0xFF5B6978),
//                             ),
//                           ),
//                         )
//                       ],
//                     )
//                   : Center(
//                       child: Text('Login success'),
//                     ),
//             ),
//           );
//   }
// }
