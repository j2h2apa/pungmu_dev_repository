import 'package:flutter/material.dart';
import 'package:audioplayers/audio_cache.dart';

void main() => runApp(XylophoneApp());

class XylophoneApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        backgroundColor: Colors.black,
        body: SafeArea(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              playSound(soundNumber: 1, color: Colors.red),
              playSound(soundNumber: 2, color: Colors.orange),
              playSound(soundNumber: 3, color: Colors.yellow),
              playSound(soundNumber: 4, color: Colors.green),
              playSound(soundNumber: 5, color: Colors.teal),
              playSound(soundNumber: 6, color: Colors.blue),
              playSound(soundNumber: 7, color: Colors.purple),
            ],
          ),
        ),
      ),
    );
  }

  Widget playSound({int soundNumber, dynamic color}) {
    return Expanded(
      child: FlatButton(
        child: null,
        color: color,
        onPressed: () {
          final player = AudioCache();
          player.play('note$soundNumber.wav');
        },
      ),
    );
  }
}
