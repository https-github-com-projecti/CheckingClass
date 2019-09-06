import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'main.dart';

class ClassListView {

  Card getStructuredGridCell(name,image,BuildContext context){
    ScreenUtil.instance = ScreenUtil.getInstance()..init(context);
    
    return Card( 
      elevation: 2.0,
      child: Column(  
        crossAxisAlignment: CrossAxisAlignment.stretch,
        mainAxisSize: MainAxisSize.max,
        verticalDirection: VerticalDirection.down,
        children: <Widget>[ 
          
        InkWell(
           
            onTap: () {
              print("tapped");
              Navigator.of(context).pushNamed("/" + homeclass);
            },
            child: SingleChildScrollView(
              child: Column(  
                crossAxisAlignment: CrossAxisAlignment.center,
                children: <Widget>[ 
                  Padding( 
                    padding: EdgeInsets.all(0.0),
                    child: Container( 
                      color: Colors.amber,
                      width: double.maxFinite,
                      height: 166.0,

                      child: Column( 
                        crossAxisAlignment: CrossAxisAlignment.center,
                        children: <Widget>[
                          Padding(
                            padding: const EdgeInsets.all(10.0),
                            child: Container( 
                              child: Column( 
                                children: <Widget>[ 
                                  Text(name),
                                  FlutterLogo(),
                                ],
                              ),
                            ),
                          )
                        ],
                      ),
                    ),
                  ),
                ],
              ),
            )
        ),
      ],
      ),
    );
  }

 GridView build(BuildContext context) {
    return GridView.count(
      primary: true,
      padding: const EdgeInsets.all(2.0),
      crossAxisCount: 3,
      childAspectRatio: 0.85,
      mainAxisSpacing: 1.0,
      crossAxisSpacing: 1.0,
      children: <Widget>[
        getStructuredGridCell("CPE", "",context),
        getStructuredGridCell("OS", "",context),
        getStructuredGridCell("Software Engineering", "",context),
        // getStructuredGridCell("Instagram", "social/instagram.png"),
        // getStructuredGridCell("Linkedin", "social/linkedin.png"),
        // getStructuredGridCell("Google Plus", "social/google_plus.png"),
        // getStructuredGridCell("Launcher Icon", "ic_launcher.png"),
      ],
    );
  }
}

  


