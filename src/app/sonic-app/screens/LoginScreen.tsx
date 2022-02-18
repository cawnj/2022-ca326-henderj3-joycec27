import { StyleSheet } from "react-native";

import { Text, View } from "../components/Themed";

export default function LoginScreen() {
  return (
    <View style={styles.container}>
      <Text>Hello from Login Page</Text>
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
});
