import 'dart:isolate';
import 'dart:math';

import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class InifiniteProcessScreen extends StatelessWidget {
  static const id = 'infinite_process_screen';
  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (context) => InfiniteProcessIsolateController(),
      child: Center(
        child: Text('infinite_process_screen'),
      ),
    );
  }
}

Future<void> _secondIsolateEntryPoint(SendPort callerSP) async {
  int multiplyValue = 1;

  ReceivePort newIceRP = ReceivePort();
  callerSP.send(newIceRP.sendPort);

  newIceRP.listen((dynamic message) {
    if (message is int) {
      multiplyValue = message;
    }

    print('come in $message');
  });

  while (true) {
    int sum = 0;
    for (int i = 0; i < 10000; i++) {
      sum += await doSomeWork();
    }

    callerSP.send(sum * multiplyValue);
  }
}

class InfiniteProcessIsolateController extends ChangeNotifier {
  Isolate newIsolate;
  ReceivePort receivePort;
  SendPort newIceSP;
  Capability capability;

  int _currentMultiplier = 1;
}

Future<int> doSomeWork() {
  var rng = Random();

  return Future<int>(() {
    int sum = 0;

    for (int i = 0; i < 1000; i++) {
      sum += rng.nextInt(100);
    }

    return sum;
  });
}
