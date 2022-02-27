import Constants from "expo-constants";
import React, { useEffect, useRef, useState } from "react";
import { Platform, StyleSheet, TouchableOpacity } from "react-native";
import * as Notifications from "expo-notifications";
import { Text, View } from "../components/Themed";
import { RootTabScreenProps } from "../types";
import { auth } from "../firebase";
import { Card, ThemeProvider } from "react-native-elements";
import useColorScheme from "../hooks/useColorScheme";

Notifications.setNotificationHandler({
  handleNotification: async () => ({
    shouldShowAlert: true,
    shouldPlaySound: true,
    shouldSetBadge: true,
  }),
});

export default function TabOneScreen({
  navigation,
}: RootTabScreenProps<"TabOne">) {
  const [expoPushToken, setExpoPushToken] = useState("");
  const [notification, setNotification] = useState(false);
  const [locationDataString, setLocationDataString] = useState("");
  const notificationListener = useRef();
  const responseListener = useRef();

  const fetchLocationData = async () => {
    const firebaseUID = auth.currentUser?.uid;
    const resp = await fetch("https://sonic.cawnj.dev/latestlocation", {
      method: "POST",
      body: JSON.stringify({
        user_id: firebaseUID,
      }),
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
    });
    const data = await resp.json();

    let locationData: string;
    // set locationDataString if a latest location exists
    if (data.name) {
      const locationName = data.name;
      const prettyDate = data.timestamp
        .split("T")[0]
        .split("-")
        .reverse()
        .join("-");
      const prettyTime = data.timestamp.split("T")[1].substring(0, 5);
      locationData = `You visited ${locationName} on ${prettyDate} at ${prettyTime}`;
    } else {
      // else set it to this
      locationData = "You have not visited anywhere yet!";
    }
    setLocationDataString(locationData);
  };

  useEffect(() => {
    registerForPushNotificationsAsync().then((token) =>
      setExpoPushToken(token)
    );

    // This listener is fired whenever a notification is received while the app is foregrounded
    notificationListener.current =
      Notifications.addNotificationReceivedListener((notification) => {
        setNotification(notification);
      });

    // This listener is fired whenever a user taps on or interacts with a notification (works when app is foregrounded, backgrounded, or killed)
    responseListener.current =
      Notifications.addNotificationResponseReceivedListener((response) => {
        const {
          notification: {
            request: {
              content: {
                data: { screen },
              },
            },
          },
        } = response;
        //when the user taps on the notification, this line checks if they //are suppose to be taken to a particular screen
        if (screen) {
          props.navigation.navigate(screen);
        }
      });

    fetchLocationData();
    return () => {
      Notifications.removeNotificationSubscription(
        notificationListener.current
      );
      Notifications.removeNotificationSubscription(responseListener.current);
    };
  }, []);

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Home</Text>
      <View
        style={styles.separator}
        lightColor="#eee"
        darkColor="rgba(255,255,255,0.1)"
      />
      <ThemeProvider useDark={useColorScheme() === "dark"}>
        <Card>
          <Card.Title>
            <Text style={styles.cardTitle}>Your Latest Visit</Text>
          </Card.Title>
          <Card.Divider />
          <Text style={styles.content}>{locationDataString}</Text>
        </Card>
      </ThemeProvider>
      <TouchableOpacity
        onPress={() => navigation.navigate("Help")}
        style={styles.button}
      >
        <Text style={styles.buttonText}>I Have Covid</Text>
      </TouchableOpacity>
    </View>
  );
}

async function registerForPushNotificationsAsync() {
  let token: string | undefined;
  if (Constants.isDevice) {
    const { status: existingStatus } =
      await Notifications.getPermissionsAsync();
    let finalStatus = existingStatus;
    if (existingStatus !== "granted") {
      const { status } = await Notifications.requestPermissionsAsync();
      finalStatus = status;
      console.log("existingStatus", existingStatus);
    }
    if (finalStatus !== "granted") {
      alert("Failed to get push token for push notification!");
      console.log("finalStatus", finalStatus);
      return;
    }
    token = (await Notifications.getExpoPushTokenAsync()).data;
    registerWithBackend(token);
  } else {
    alert("Must use physical device for Push Notifications");
  }

  if (Platform.OS === "android") {
    await Notifications.setNotificationChannelAsync("default", {
      name: "default",
      showBadge: true,
      importance: Notifications.AndroidImportance.MAX,
      vibrationPattern: [0, 250, 250, 250],
      lightColor: "#FE9018",
    });
  }

  return token;
}

async function registerWithBackend(token: string) {
  const firebaseUID = auth.currentUser?.uid;
  console.log(firebaseUID);
  console.log(token);

  fetch("https://sonic.cawnj.dev/register", {
    method: "POST",
    body: JSON.stringify({
      user_id: firebaseUID,
      expo_token: token,
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
  content: {
    fontSize: 12,
    fontWeight: "bold",
  },
  cardTitle: {
    fontSize: 20,
    fontWeight: "bold",
  },
  separator: {
    marginVertical: 30,
    height: 1,
    width: "80%",
  },
  button: {
    backgroundColor: "#FF0000",
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
