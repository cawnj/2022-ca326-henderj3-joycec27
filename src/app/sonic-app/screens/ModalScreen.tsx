import { StatusBar } from "expo-status-bar";
import { Platform, StyleSheet, TouchableOpacity } from "react-native";

import { Text, View } from "../components/Themed";
import { auth } from "../firebase";
import { useNavigation } from "@react-navigation/native";

export default function ModalScreen() {
  const navigation = useNavigation();

  // send trace request via POST req
  const sendPostRequest = () => {
    const firebaseUID = auth.currentUser?.uid;
    fetch("https://sonic.cawnj.dev/trace", {
      method: "POST",
      body: JSON.stringify({
        user_id: firebaseUID,
      }),
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
    })
      .then((response) => response.json())
      .then((responseJson) => {
        console.log(responseJson);
      })
      .catch((error) => {
        console.log(error);
      });
    // return user to home page after req is executed
    navigation.navigate("Root");
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Help</Text>
      <View
        style={styles.separator}
        lightColor="#eee"
        darkColor="rgba(255,255,255,0.1)"
      />
      {/* Use a light status bar on iOS to account for the black space above the modal */}
      <Text
        style={styles.getStartedText}
        lightColor="rgba(0,0,0,0.8)"
        darkColor="rgba(255,255,255,0.8)"
      >
        If you have Covid please press
      </Text>
      <TouchableOpacity onPress={sendPostRequest} style={styles.button}>
        <Text style={styles.buttonText}>Continue</Text>
      </TouchableOpacity>
      <StatusBar style={Platform.OS === "ios" ? "light" : "auto"} />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: "center",
    justifyContent: "center",
  },
  title: {
    fontSize: 20,
    fontWeight: "bold",
  },
  separator: {
    marginVertical: 30,
    height: 1,
    width: "80%",
  },
  getStartedText: {
    fontSize: 17,
    lineHeight: 24,
    textAlign: "center",
  },
  button: {
    backgroundColor: "#26DFD0",
    width: "60%",
    padding: 15,
    borderRadius: 10,
    alignItems: "center",
    marginTop: 40,
  },
  buttonText: {
    color: "white",
    fontWeight: "700",
    fontSize: 16,
  },
});
