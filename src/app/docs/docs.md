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
- `yarn global add expo-cli`
- `export PATH=~/.yarn/bin:$PATH`

### Create Project - Choose one
- `npx react-native init sonic_app --template react-native-template-typescript`
- `expo init sonic_app`
