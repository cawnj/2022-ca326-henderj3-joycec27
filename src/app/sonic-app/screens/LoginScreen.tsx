import { StyleSheet } from "react-native";

import {
  Text,
  TextInput,
  View,
  KeyboardAvoidingView,
  TouchableOpacity,
} from "../components/Themed";

export default function LoginScreen() {
  return (
    <KeyboardAvoidingView style={styles.container}>
      <View style={styles.container}>
        <TextInput
          placeholder="Email"
          // value={ }
          // onChangeText={text => ""}
          style={styles.input}
        />
        <TextInput
          placeholder="Password"
          // value={ }
          // onChangeText={text => ""}
          style={styles.input}
          secureTextEntry
        />
      </View>
      <View style={styles.container}>
        <TouchableOpacity onPress={() => {}} style={styles.button}>
          <Text style={styles.buttonText}>Login</Text>
        </TouchableOpacity>
        <TouchableOpacity
          onPress={() => { }}
          style={[styles.button, styles.buttonOutline]}
        >
          <Text style={[styles.buttonText, styles.buttonOutlineText]}>
            Register
          </Text>
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
  input: {},
  button: {
  },
  buttonText: {
  },
  buttonOutline: {
  },
  buttonOutlineText: {
  },
});
