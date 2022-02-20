// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: user.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

public struct V1_MakeUser {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var token: String = String()

  public var user: V1_User {
    get {return _user ?? V1_User()}
    set {_user = newValue}
  }
  /// Returns true if `user` has been explicitly set.
  public var hasUser: Bool {return self._user != nil}
  /// Clears the value of `user`. Subsequent reads from it will return its default value.
  public mutating func clearUser() {self._user = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _user: V1_User? = nil
}

public struct V1_User {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var userID: String = String()

  /// required encrypt
  public var userPassword: String = String()

  public var userNickname: String = String()

  public var userEmail: String {
    get {return _userEmail ?? String()}
    set {_userEmail = newValue}
  }
  /// Returns true if `userEmail` has been explicitly set.
  public var hasUserEmail: Bool {return self._userEmail != nil}
  /// Clears the value of `userEmail`. Subsequent reads from it will return its default value.
  public mutating func clearUserEmail() {self._userEmail = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _userEmail: String? = nil
}

public struct V1_UserDetail {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var id: String = String()

  public var userID: String = String()

  public var userNickname: String = String()

  public var userEmail: String {
    get {return _userEmail ?? String()}
    set {_userEmail = newValue}
  }
  /// Returns true if `userEmail` has been explicitly set.
  public var hasUserEmail: Bool {return self._userEmail != nil}
  /// Clears the value of `userEmail`. Subsequent reads from it will return its default value.
  public mutating func clearUserEmail() {self._userEmail = nil}

  public var kubeNamespace: [V1_Namespace] = []

  public var userOranization: [V1_Organization] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _userEmail: String? = nil
}

public struct V1_MakeOrganization {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var token: String = String()

  public var organization: V1_Organization {
    get {return _organization ?? V1_Organization()}
    set {_organization = newValue}
  }
  /// Returns true if `organization` has been explicitly set.
  public var hasOrganization: Bool {return self._organization != nil}
  /// Clears the value of `organization`. Subsequent reads from it will return its default value.
  public mutating func clearOrganization() {self._organization = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _organization: V1_Organization? = nil
}

public struct V1_Organization {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var id: V1_Uuid {
    get {return _id ?? V1_Uuid()}
    set {_id = newValue}
  }
  /// Returns true if `id` has been explicitly set.
  public var hasID: Bool {return self._id != nil}
  /// Clears the value of `id`. Subsequent reads from it will return its default value.
  public mutating func clearID() {self._id = nil}

  /// required encrypt
  public var userNickname: String = String()

  public var userEmail: String {
    get {return _userEmail ?? String()}
    set {_userEmail = newValue}
  }
  /// Returns true if `userEmail` has been explicitly set.
  public var hasUserEmail: Bool {return self._userEmail != nil}
  /// Clears the value of `userEmail`. Subsequent reads from it will return its default value.
  public mutating func clearUserEmail() {self._userEmail = nil}

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _id: V1_Uuid? = nil
  fileprivate var _userEmail: String? = nil
}

public struct V1_OrganizationDetail {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var id: V1_Uuid {
    get {return _id ?? V1_Uuid()}
    set {_id = newValue}
  }
  /// Returns true if `id` has been explicitly set.
  public var hasID: Bool {return self._id != nil}
  /// Clears the value of `id`. Subsequent reads from it will return its default value.
  public mutating func clearID() {self._id = nil}

  /// required encrypt
  public var userNickname: String = String()

  public var userEmail: String {
    get {return _userEmail ?? String()}
    set {_userEmail = newValue}
  }
  /// Returns true if `userEmail` has been explicitly set.
  public var hasUserEmail: Bool {return self._userEmail != nil}
  /// Clears the value of `userEmail`. Subsequent reads from it will return its default value.
  public mutating func clearUserEmail() {self._userEmail = nil}

  public var user: [V1_UserDetail] = []

  public var kubeNamespace: [V1_Namespace] = []

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}

  fileprivate var _id: V1_Uuid? = nil
  fileprivate var _userEmail: String? = nil
}

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "v1"

extension V1_MakeUser: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".MakeUser"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "token"),
    2: .same(proto: "user"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.token) }()
      case 2: try { try decoder.decodeSingularMessageField(value: &self._user) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.token.isEmpty {
      try visitor.visitSingularStringField(value: self.token, fieldNumber: 1)
    }
    try { if let v = self._user {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: V1_MakeUser, rhs: V1_MakeUser) -> Bool {
    if lhs.token != rhs.token {return false}
    if lhs._user != rhs._user {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension V1_User: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".User"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "user_id"),
    2: .standard(proto: "user_password"),
    3: .standard(proto: "user_nickname"),
    4: .standard(proto: "user_email"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.userID) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.userPassword) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.userNickname) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self._userEmail) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.userID.isEmpty {
      try visitor.visitSingularStringField(value: self.userID, fieldNumber: 1)
    }
    if !self.userPassword.isEmpty {
      try visitor.visitSingularStringField(value: self.userPassword, fieldNumber: 2)
    }
    if !self.userNickname.isEmpty {
      try visitor.visitSingularStringField(value: self.userNickname, fieldNumber: 3)
    }
    try { if let v = self._userEmail {
      try visitor.visitSingularStringField(value: v, fieldNumber: 4)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: V1_User, rhs: V1_User) -> Bool {
    if lhs.userID != rhs.userID {return false}
    if lhs.userPassword != rhs.userPassword {return false}
    if lhs.userNickname != rhs.userNickname {return false}
    if lhs._userEmail != rhs._userEmail {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension V1_UserDetail: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".UserDetail"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "id"),
    2: .standard(proto: "user_id"),
    3: .standard(proto: "user_nickname"),
    4: .standard(proto: "user_email"),
    5: .standard(proto: "kube_namespace"),
    6: .standard(proto: "user_oranization"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.id) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.userID) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self.userNickname) }()
      case 4: try { try decoder.decodeSingularStringField(value: &self._userEmail) }()
      case 5: try { try decoder.decodeRepeatedMessageField(value: &self.kubeNamespace) }()
      case 6: try { try decoder.decodeRepeatedMessageField(value: &self.userOranization) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.id.isEmpty {
      try visitor.visitSingularStringField(value: self.id, fieldNumber: 1)
    }
    if !self.userID.isEmpty {
      try visitor.visitSingularStringField(value: self.userID, fieldNumber: 2)
    }
    if !self.userNickname.isEmpty {
      try visitor.visitSingularStringField(value: self.userNickname, fieldNumber: 3)
    }
    try { if let v = self._userEmail {
      try visitor.visitSingularStringField(value: v, fieldNumber: 4)
    } }()
    if !self.kubeNamespace.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.kubeNamespace, fieldNumber: 5)
    }
    if !self.userOranization.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.userOranization, fieldNumber: 6)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: V1_UserDetail, rhs: V1_UserDetail) -> Bool {
    if lhs.id != rhs.id {return false}
    if lhs.userID != rhs.userID {return false}
    if lhs.userNickname != rhs.userNickname {return false}
    if lhs._userEmail != rhs._userEmail {return false}
    if lhs.kubeNamespace != rhs.kubeNamespace {return false}
    if lhs.userOranization != rhs.userOranization {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension V1_MakeOrganization: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".MakeOrganization"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "token"),
    2: .same(proto: "organization"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.token) }()
      case 2: try { try decoder.decodeSingularMessageField(value: &self._organization) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    if !self.token.isEmpty {
      try visitor.visitSingularStringField(value: self.token, fieldNumber: 1)
    }
    try { if let v = self._organization {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 2)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: V1_MakeOrganization, rhs: V1_MakeOrganization) -> Bool {
    if lhs.token != rhs.token {return false}
    if lhs._organization != rhs._organization {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension V1_Organization: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".Organization"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "id"),
    2: .standard(proto: "user_nickname"),
    3: .standard(proto: "user_email"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._id) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.userNickname) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self._userEmail) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._id {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    if !self.userNickname.isEmpty {
      try visitor.visitSingularStringField(value: self.userNickname, fieldNumber: 2)
    }
    try { if let v = self._userEmail {
      try visitor.visitSingularStringField(value: v, fieldNumber: 3)
    } }()
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: V1_Organization, rhs: V1_Organization) -> Bool {
    if lhs._id != rhs._id {return false}
    if lhs.userNickname != rhs.userNickname {return false}
    if lhs._userEmail != rhs._userEmail {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension V1_OrganizationDetail: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".OrganizationDetail"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "id"),
    2: .standard(proto: "user_nickname"),
    3: .standard(proto: "user_email"),
    4: .same(proto: "user"),
    5: .standard(proto: "kube_namespace"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularMessageField(value: &self._id) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.userNickname) }()
      case 3: try { try decoder.decodeSingularStringField(value: &self._userEmail) }()
      case 4: try { try decoder.decodeRepeatedMessageField(value: &self.user) }()
      case 5: try { try decoder.decodeRepeatedMessageField(value: &self.kubeNamespace) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every if/case branch local when no optimizations
    // are enabled. https://github.com/apple/swift-protobuf/issues/1034 and
    // https://github.com/apple/swift-protobuf/issues/1182
    try { if let v = self._id {
      try visitor.visitSingularMessageField(value: v, fieldNumber: 1)
    } }()
    if !self.userNickname.isEmpty {
      try visitor.visitSingularStringField(value: self.userNickname, fieldNumber: 2)
    }
    try { if let v = self._userEmail {
      try visitor.visitSingularStringField(value: v, fieldNumber: 3)
    } }()
    if !self.user.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.user, fieldNumber: 4)
    }
    if !self.kubeNamespace.isEmpty {
      try visitor.visitRepeatedMessageField(value: self.kubeNamespace, fieldNumber: 5)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: V1_OrganizationDetail, rhs: V1_OrganizationDetail) -> Bool {
    if lhs._id != rhs._id {return false}
    if lhs.userNickname != rhs.userNickname {return false}
    if lhs._userEmail != rhs._userEmail {return false}
    if lhs.user != rhs.user {return false}
    if lhs.kubeNamespace != rhs.kubeNamespace {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}
