import { useEffect, useState } from "react";
import { StyleSheet } from "react-native";
import { useNavigation } from "@react-navigation/native";
import { MaterialIcons } from "@expo/vector-icons";
import { auth } from "../firebase";

import {
  Text,
  TextInput,
  View,
  KeyboardAvoidingView,
  TouchableOpacity,
} from "../components/Themed";

export default function ForgotPasswordScreen() {
  const [email, setEmail] = useState("");

  const navigation = useNavigation();

  const resetPassword = () => {
    auth.sendPasswordResetEmail(email).then(
      () => {
        alert("Password reset email has been sent.");
        navigation.navigate("Login");
      },
      (error) => {
        alert(error.message);
      }
    );
  };

  return (
    <KeyboardAvoidingView style={styles.container}>
      <View style={styles.inputContainer}>
        <TextInput
          placeholder="Email"
          value={email}
          onChangeText={(text) => setEmail(text)}
          style={styles.input}
        />
      </View>
      <View style={[styles.container, styles.buttonContainer]}>
        <TouchableOpacity onPress={resetPassword} style={styles.button}>
          <Text style={styles.buttonText}>Reset Password</Text>
        </TouchableOpacity>
        <TouchableOpacity
          onPress={() => navigation.navigate("Login")}
          style={styles.Bbutton}
        >
          <Text style={styles.BbuttonText}>Back</Text>
        </TouchableOpacity>
      </View>
    </KeyboardAvoidingView>
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
  inputContainer: {
    width: "75%",
    marginTop: 100,
  },
  input: {
    paddingHorizontal: 15,
    paddingVertical: 10,
    borderRadius: 10,
    marginTop: 5,
  },
  buttonContainer: {
    width: "60%",
    marginTop: 40,
  },
  button: {
    backgroundColor: "#0782F9",
    width: "100%",
    padding: 15,
    borderRadius: 10,
    alignItems: "center",
  },
  buttonText: {
    color: "white",
    fontWeight: "700",
    fontSize: 16,
  },
  Bbutton: {
    width: "100%",
    padding: 15,
    borderRadius: 10,
    alignItems: "center",
  },
  BbuttonText: {
    color: "#0782F9",
    fontWeight: "700",
    fontSize: 16,
  },
});
