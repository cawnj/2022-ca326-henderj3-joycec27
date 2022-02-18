// Import the functions you need from the SDKs you need
import firebase from "firebase/compat/app";
import "firebase/compat/auth";
import "firebase/compat/firestore";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyCHb34Qxrj_lh3XU9ARTnew_uImtDeuVso",
  authDomain: "fir-auth-451b3.firebaseapp.com",
  projectId: "fir-auth-451b3",
  storageBucket: "fir-auth-451b3.appspot.com",
  messagingSenderId: "854580616381",
  appId: "1:854580616381:web:64e493da948fd3efdb6bd4",
};

// Initialize Firebase
let app;
if (firebase.apps.length == 0) {
  app = firebase.initializeApp(firebaseConfig);
} else {
  app = firebase.app();
}

const auth = firebase.auth();

export { auth };
