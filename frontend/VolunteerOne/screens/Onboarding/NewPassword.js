import React from "react";
// import zxcvbn from "zxcvbn"; // for password strength
import { useState } from "react";
import {
  StyleSheet,
  ImageBackground,
  Dimensions,
  StatusBar,
  KeyboardAvoidingView,
  Image,
  TextInput,
  TouchableOpacity
} from "react-native";
import { Block, Text } from "galio-framework";

import { Button } from "../../components";
import { Images, argonTheme } from "../../constants";

const { width, height } = Dimensions.get("screen");

/** ==================================== New Password Screen ==================================== **/

const NewPassword = ({ navigation }) => {
  const [newPassword, setNewPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [passwordStrength, setPasswordStrength] = useState('');

  const handleNewPasswordInput = (input) => {
    setNewPassword(input);
    setPasswordStrength(checkPasswordStrength(input));
  };

  const handleConfirmPasswordInput = (input) => {
    setConfirmPassword(input);
  };

  const handleSaveBtnClick = () => {
    console.log(newPassword, confirmPassword);
    navigation.navigate('Login');
  };

  const checkPasswordStrength = (password) => {
    if (password.length < 8) {
      return "Weak";
    }
    if (password.length < 12) {
      return "Moderate";
    }
    return "Strong";
  };

  return (
    <Block flex middle>
      <StatusBar hidden />
      <ImageBackground
        source={Images.Onboarding}
        style={{ width, height, zIndex: 1 }}
      >
        <Block safe flex middle>
          <Block style={styles.loginContainer}>
            <Block flex>
              <Block flex={0.5} middle style={styles.instructionText}>
                <Block center>
                  <Image source={Images.VolunteerOneIcon} style={styles.logo} />
                </Block>
              </Block>
              <Block flex center>
                <KeyboardAvoidingView
                  style={{ flex: 1 }}
                  behavior="padding"
                  enabled
                >
                  <Block flex={0.2} middle style={styles.instructionText}>
              <Text
                    color="#8898AA"
                    size={12}
                    style={{
                      fontWeight: "bold",
                    }}
                  >
                    Enter your new passoword
                  </Text>
              </Block>
                  <Block width={width * 0.8} style={{ marginBottom: 15 }}>
                    <TextInput
                      secureTextEntry={true}
                      style={styles.input}
                      placeholder="New Password"
                      onChangeText={handleNewPasswordInput}
                    />
                  </Block>
                  
                  <Block width={width * 0.8}>
                    <TextInput
                      secureTextEntry={true}
                      style={styles.input}
                      placeholder="Comfirm Password"
                      onChangeText={handleConfirmPasswordInput}
                    />
                  </Block>
                  <Block middle>
                  <Text
                      color={argonTheme.COLORS.PRIMARY}
                      size={14}
                      style={{
                        fontWeight: "bold",
                        marginTop: 20,
                      }}
                    >
                      Password strength: {passwordStrength}
                    </Text>
                    <Button
                      color="primary"
                      style={styles.createButton}
                      onPress={handleSaveBtnClick}
                    >
                      <Text bold size={14} color={argonTheme.COLORS.WHITE}>
                        CREATE PASSWORD
                      </Text>
                    </Button>
                  </Block>
                </KeyboardAvoidingView>
              </Block>
            </Block>
          </Block>
        </Block>
      </ImageBackground>
    </Block>
  );
};
// }

const styles = StyleSheet.create({
  input: {
    borderColor: argonTheme.COLORS.BORDER,
    height: 44,
    backgroundColor: "#FFFFFF",
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: { width: 0, height: 1 },
    shadowRadius: 2,
    shadowOpacity: 0.05,
    elevation: 2,
    paddingLeft: 10,
  },
  instructionText: {
    flexDirection: "row",
  },
  loginContainer: {
    width: width * 0.9,
    height: height * 0.75,
    backgroundColor: "#F4F5F7",
    borderRadius: 4,
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: {
      width: 0,
      height: 4,
    },
    shadowRadius: 8,
    shadowOpacity: 0.1,
    elevation: 1,
    overflow: "hidden",
  },
  inputIcons: {
    marginRight: 12,
  },
  passwordCheck: {
    paddingLeft: 15,
    paddingTop: 13,
    paddingBottom: 30,
  },
  createButton: {
    width: width * 0.5,
    marginTop: 25,
  },
  logo: {
    width: 265,
    height: 50,
    zIndex: 2,
    position: 'relative',
    marginTop: '20%'
  },
});

export default NewPassword;
