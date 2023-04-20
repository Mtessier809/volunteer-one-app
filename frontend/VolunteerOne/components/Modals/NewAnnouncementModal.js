import React from "react";
import {
  Alert,
  Modal,
  StyleSheet,
  Pressable,
  View,
  Dimensions,
  TextInput,
  Image,
} from "react-native";
import { Block, Text, theme } from "galio-framework";
import { argonTheme } from "../../constants";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";

const { width, height } = Dimensions.get("screen");

/** ==================================== New Announcement Modal Component ==================================== **/

class NewAnnouncementModal extends React.Component {
  state = {
    datetime: new Date(),
    title: "",
    description: "",
  };

  render() {
    // const handleAddNewClick = () => {
    //   console.log("Adding New Announcement ",this.state)
    //   // post to db
    // }

    return (
      <View style={styles.centeredView}>
        <Modal
          animationType="fade"
          transparent={true}
          visible={this.props.visible}
          onRequestClose={() => {
            console.log("Modal has been closed.");
            this.props.handleModalVisible();
          }}
        >
          <View style={[styles.centeredView, styles.modalViewOutside]}>
            <View style={styles.modalView}>
              {/* exit modal */}
              <Pressable
                onPress={() => this.props.handleModalVisible()}
                style={{ alignItems: "flex-end", margin: 5 }}
              >
                <MaterialCommunityIcons
                  size={24}
                  name="close"
                  color={theme.COLORS.ICON}
                />
              </Pressable>

              <View style={styles.modalViewInside}>
                <Text style={styles.header}>New Announcement</Text>

                <Text style={styles.secondaryHeader}>Post title</Text>
                <Block width={width * 0.8 - 20} style={{ marginBottom: 15 }}>
                  <TextInput
                    style={styles.input}
                    placeholder="Enter a title"
                    placeholderTextColor={"lightgrey"}
                    onChangeText={(e) => this.setState({ title: e })}
                  />
                </Block>

                <Text style={styles.secondaryHeader}>Description</Text>
                <Block width={width * 0.8 - 20} style={{ marginBottom: 15 }}>
                  <TextInput
                    style={styles.input}
                    placeholder="Provide announcement details here"
                    placeholderTextColor={"lightgrey"}
                    height={height * 0.3}
                    textAlignVertical={"top"}
                    paddingTop={10}
                    multiline={true}
                    onChangeText={(e) => this.setState({ description: e })}
                  />
                </Block>

                <Pressable
                  style={[styles.button, styles.buttonClose]}
                  onPress={() => {
                    this.props.handleModalVisible();
                    this.props.addNewAnnouncement(this.state);
                  }}
                >
                  <Text style={styles.textStyle}>CREATE ANNOUNCEMENT</Text>
                </Pressable>
              </View>
            </View>
          </View>
        </Modal>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  header: {
    fontSize: 25,
    fontWeight: "bold",
    color: "#525F7F",
    marginBottom: 20,
  },
  secondaryHeader: {
    fontSize: 17,
    fontWeight: "bold",
    color: "#525F7F",
    marginBottom: 5,
  },
  input: {
    borderColor: argonTheme.COLORS.BORDER,
    borderWidth: 0.5,
    borderRadius: 5,
    height: 44,
    backgroundColor: "#FFFFFF",
    shadowColor: argonTheme.COLORS.BLACK,
    shadowOffset: { width: 0, height: 1 },
    shadowRadius: 2,
    shadowOpacity: 0.05,
    elevation: 2,
    paddingLeft: 10,
  },
  modalView: {
    backgroundColor: "white",
    shadowColor: "#000",
    shadowOffset: {
      width: 5,
      height: 5,
    },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 5,
  },

  // matt's added styles above ^^^

  centeredView: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    // marginTop: 22,
  },
  modalViewOutside: {
    backgroundColor: "rgba(52, 52, 52, 0.75)", // changed opacity of background when modal is open
  },
  modalViewInside: {
    padding: 25,
    paddingTop: 0,
  },
  button: {
    borderRadius: 5,
    padding: 10,
    elevation: 2,
  },
  buttonClose: {
    backgroundColor: "#5e72e4",
    padding: 10,
    marginTop: 10,
  },
  textStyle: {
    color: "white",
    fontWeight: "bold",
    textAlign: "center",
  },
});

export default NewAnnouncementModal;
