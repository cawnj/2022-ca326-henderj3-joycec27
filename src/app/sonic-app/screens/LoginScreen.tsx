import { StyleSheet } from "react-native";

import { TextInput, View, KeyboardAvoidingView } from "../components/Themed";

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
  input: {
  },
});
