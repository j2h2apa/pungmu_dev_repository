import 'package:async/async.dart';
import 'package:flutter/material.dart';
import 'package:flutter/foundation.dart';

int fib(int n) {
  var a = n - 1;
  var b = n - 2;

  if (n == 1) {
    return 0;
  } else if (n == 0) {
    return 1;
  } else {
    return (fib(a) + fib(b));
  }
}

class PerformanceScreen extends StatefulWidget {
  static const id = 'performance_screen';

  @override
  _PerformanceScreenState createState() => _PerformanceScreenState();
}

class _PerformanceScreenState extends State<PerformanceScreen> {
  Future<void> computeFuture = Future.value();
  final AsyncMemoizer _memoizer = AsyncMemoizer();

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        children: [
          SmoothAnimationWidget(),
          Container(
            alignment: Alignment.bottomCenter,
            padding: EdgeInsets.only(top: 150.0),
            child: Column(
              children: [
                FutureBuilder(
                  future: computeFuture,
                  builder: (context, snapshot) {
                    return RaisedButton(
                        child: Text('Compute on Main'),
                        elevation: 8.0,
                        onPressed: () {
                          if (snapshot.connectionState ==
                              ConnectionState.done) {
                            handleComputeOnMain(context);
                          } else if (snapshot.connectionState ==
                              ConnectionState.waiting) {
                            return Center(child: CircularProgressIndicator());
                          } else if (snapshot.connectionState ==
                              ConnectionState.active) {}

                          if (snapshot.hasError) {
                            print('shapshot error!');
                          }
                        });
                  },
                ),
                FutureBuilder(
                  future: computeFuture,
                  builder: (context, snapshot) {
                    if (snapshot.connectionState != ConnectionState.done) {
                      return Center(
                          child: CircularProgressIndicator(
                        backgroundColor: Colors.black,
                      ));
                    }

                    return RaisedButton(
                      child: Text('Compute on Secondary'),
                      elevation: 8.0,
                      onPressed: () {
                        if (snapshot.connectionState == ConnectionState.done) {
                          handleComputeOnSecondary(context);
                        } else {
                          print('tttttttttttttt');
                        }
                      },
                    );
                  },
                ),
                Center(
                    child: CircularProgressIndicator(
                  backgroundColor: Colors.black,
                )),
              ],
            ),
          ),
        ],
      ),
    );
  }

  void handleComputeOnMain(BuildContext context) {
    var future = computeOnMainIsolate()
      ..then((_) {
        var snackBar = SnackBar(
          content: Text('Main Isolate Done!'),
        );
        Scaffold.of(context).showSnackBar(snackBar);
      });

    setState(() {
      print('setState');
      computeFuture = future;
    });
  }

  void handleComputeOnSecondary(BuildContext context) {
    var future = computeOnSecondaryIsolate()
      ..then((_) {
        var snackBar = SnackBar(
          content: Text('Secondary Isolate Done!'),
        );
        Scaffold.of(context).showSnackBar(snackBar);
      });

    setState(() {
      computeFuture = future;
    });
  }

  Future<void> computeOnMainIsolate() async {
    // return this._memoizer.runOnce(() async {
    //   await Future<void>.delayed(Duration(milliseconds: 100));
    //   fib(45);
    // });

    await Future<void>.delayed(Duration(milliseconds: 1));
    print('Result : ' + fib(45).toString());
  }

  Future<void> computeOnSecondaryIsolate() async {
    await compute(fib, 45);
  }
}

class SmoothAnimationWidget extends StatefulWidget {
  @override
  _SmoothAnimationWidgetState createState() => _SmoothAnimationWidgetState();
}

class _SmoothAnimationWidgetState extends State<SmoothAnimationWidget>
    with TickerProviderStateMixin {
  AnimationController _animationController;
  Animation<BorderRadius> _borderAnimation;

  @override
  void initState() {
    super.initState();

    _animationController = AnimationController(
      duration: const Duration(seconds: 1),
      vsync: this,
      upperBound: 1.0,
    );

    _borderAnimation = BorderRadiusTween(
      begin: BorderRadius.circular(100.0),
      end: BorderRadius.circular(0.0),
    ).animate(_animationController);

    _borderAnimation.addStatusListener((AnimationStatus state) {
      // print(state);
    });

    _animationController.repeat(reverse: true);
  }

  @override
  void dispose() {
    _animationController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: AnimatedBuilder(
        animation: _borderAnimation,
        builder: (context, child) {
          return Container(
            child: FlutterLogo(
              size: 200.0,
            ),
            alignment: Alignment.bottomCenter,
            width: 300.0,
            height: 200.0,
            decoration: BoxDecoration(
              gradient: LinearGradient(begin: Alignment.topLeft, colors: [
                Colors.blueAccent,
                Colors.redAccent,
              ]),
              borderRadius: _borderAnimation.value,
            ),
          );
        },
      ),
    );
  }
}
