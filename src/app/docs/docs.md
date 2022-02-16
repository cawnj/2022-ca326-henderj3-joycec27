# Sonic App

## Documentation/Help

### Node Setup
- `brew install nvm`
- `nvm install lts/fermium`
- `nvm use lts/fermium`
- `nvm alias default lts/fermium`

### Dependency Setup
- `brew install watchman ruby`
- `export PATH=/opt/homebrew/opt/ruby/bin:$PATH`
- ```export PATH=`gem environment gemdir`/bin:$PATH```
- `gem install ffi cocoapods`
- `brew install --cask adoptopenjdk/openjdk/adoptopenjdk8`
- Install Xcode
  - Preferences, Locations, Command Line Tools, click on version
  - Preferences, Components, select latest iOS simulator

### Expo Setup
- `yarn global add expo expo-cli`
- `export PATH=~/.yarn/bin:$PATH`

### Create Project - Choose one
- `npx react-native init sonic-app --template react-native-template-typescript`
- `expo init sonic-app`

### Expo Build
- expo build:ios has been superseded by eas build
- `yarn global add eas-cli`
- Build docs [here](https://docs.expo.dev/build/setup/)
- To generate eas.json: `eas build:configure`
- To build Sonic.app for iOS Simulator: `eas build -p ios --profile simulator`
- To publish to expo: `expo publish`; Run on device with "Expo Go" app

### VSCode Extensions
- React Native Tools
- React-Native/React/Redux snippets
- ESLint & Prettier

### ESLint & Prettier Setup
- `yarn add eslint --dev`
- `yarn run eslint --init`
- `yarn add eslint-plugin-react@latest @typescript-eslint/eslint-plugin@latest @typescript-eslint/parser@latest --dev`
- `yarn add eslint-config-prettier eslint-plugin-prettier prettier --dev`
- See .vscode/settings.json and .eslintrc.json for config
