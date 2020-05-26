package main

/*
Envelope
|\_ Payload []byte Payload.toBuffer
	|\_ Header *Header
		|\_ ChannelHeader   []byte
			|\_ Type int32
			|\_ Version int32
			|\_ Timestamp *google_protobuf.Timestamp
			|\_ ChannelId string
			|\_ TxId string
			|\_ Epoch uint64
			|\_ Extension []byte
			|\_ TlsCertHash []byte
		|\_ SignatureHeader []byte
	|\_ Date []byte ConfigUpdateEnvelope.toBuffer
		|\_ ConfigUpdate []byte ConfigUpdate.toBuffer
			|\_ ChannelId    string
			|\_ ReadSet      *ConfigGroup
			|\_ WriteSet     *ConfigGroup
			|\_ IsolatedData map[string][]byte
		|\_ Signatures   []*ConfigSignature
			|\_ SignatureHeader []byte
			|\_ Signature       []byte
|\_ Signature []byte

*/

type Envelope struct {
	//*// A marshaled Payload
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// A signature by the creator specified in the Payload header
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

type Block struct {
	Header   *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Data     *BlockData     `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	Metadata *BlockMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

type BlockHeader struct {
	Number       uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	PreviousHash []byte `protobuf:"bytes,2,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	DataHash     []byte `protobuf:"bytes,3,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
}

type BlockData struct {
	Data [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`   //HeaderType_name
}

type BlockMetadata struct {
	Metadata [][]byte `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

type Metadata struct {
	Value      []byte               `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Signatures []*MetadataSignature `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
}

type MetadataSignature struct {
	SignatureHeader []byte `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	Signature       []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

var HeaderType_name = map[int32]string{
	0: "MESSAGE",
	1: "CONFIG",                   //Payload
	2: "CONFIG_UPDATE",            //Config
	3: "ENDORSER_TRANSACTION",
	4: "ORDERER_TRANSACTION",
	5: "DELIVER_SEEK_INFO",
	6: "CHAINCODE_PACKAGE",
	8: "PEER_ADMIN_OPERATION",
}

type Payload struct {
	// Header is included to provide identity and prevent replay
	Header *Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// Data, the encoding of which is defined by the type in the header
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

type Header struct {
	ChannelHeader   []byte `protobuf:"bytes,1,opt,name=channel_header,json=channelHeader,proto3" json:"channel_header,omitempty"`
	SignatureHeader []byte `protobuf:"bytes,2,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
}


type ChannelHeader struct {
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// Identifier of the channel this message is bound for
	ChannelId string `protobuf:"bytes,4,opt,name=channel_id,json=channelId" json:"channel_id,omitempty"`
	// An unique identifier that is used end-to-end.
	//  -  set by higher layers such as end user or SDK
	//  -  passed to the endorser (which will check for uniqueness)
	//  -  as the header is passed along unchanged, it will be
	//     be retrieved by the committer (uniqueness check here as well)
	//  -  to be stored in the ledger
	TxId string `protobuf:"bytes,5,opt,name=tx_id,json=txId" json:"tx_id,omitempty"`
	// The epoch in which this header was generated, where epoch is defined based on block height
	// Epoch in which the response has been generated. This field identifies a
	// logical window of time. A proposal response is accepted by a peer only if
	// two conditions hold:
	// 1. the epoch specified in the message is the current epoch
	// 2. this message has been only seen once during this epoch (i.e. it hasn't
	//    been replayed)
	Epoch uint64 `protobuf:"varint,6,opt,name=epoch" json:"epoch,omitempty"`
	// Extension that may be attached based on the header type
	Extension []byte `protobuf:"bytes,7,opt,name=extension,proto3" json:"extension,omitempty"`
	// If mutual TLS is employed, this represents
	// the hash of the client's TLS certificate
	TlsCertHash []byte `protobuf:"bytes,8,opt,name=tls_cert_hash,json=tlsCertHash,proto3" json:"tls_cert_hash,omitempty"`
}

type SignatureHeader struct {
	// Creator of the message, a marshaled msp.SerializedIdentity
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Arbitrary number that may only be used once. Can be used to detect replay attacks.
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

type Config struct {
	Sequence     uint64       `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	ChannelGroup *ConfigGroup `protobuf:"bytes,2,opt,name=channel_group,json=channelGroup" json:"channel_group,omitempty"`
}

type ConfigGroup struct {
	Version   uint64                   `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Groups    map[string]*ConfigGroup  `protobuf:"bytes,2,rep,name=groups" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Values    map[string]*ConfigValue  `protobuf:"bytes,3,rep,name=values" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Policies  map[string]*ConfigPolicy `protobuf:"bytes,4,rep,name=policies" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ModPolicy string                   `protobuf:"bytes,5,opt,name=mod_policy,json=modPolicy" json:"mod_policy,omitempty"`
}

type ConfigValue struct {
	Version   uint64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Value     []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ModPolicy string `protobuf:"bytes,3,opt,name=mod_policy,json=modPolicy" json:"mod_policy,omitempty"`
}

type ConfigPolicy struct {
	Version   uint64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Policy    *Policy `protobuf:"bytes,2,opt,name=policy" json:"policy,omitempty"`
	ModPolicy string  `protobuf:"bytes,3,opt,name=mod_policy,json=modPolicy" json:"mod_policy,omitempty"`
}

type Policy struct {
	Type  int32  `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

type ConfigEnvelope struct {
	Config     *Config   `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	LastUpdate *Envelope `protobuf:"bytes,2,opt,name=last_update,json=lastUpdate" json:"last_update,omitempty"`
}

type ConfigUpdateEnvelope struct {
	ConfigUpdate []byte             `protobuf:"bytes,1,opt,name=config_update,json=configUpdate,proto3" json:"config_update,omitempty"`
	Signatures   []*ConfigSignature `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty"`
}

type ConfigUpdate struct {
	ChannelId    string            `protobuf:"bytes,1,opt,name=channel_id,json=channelId" json:"channel_id,omitempty"`
	ReadSet      *ConfigGroup      `protobuf:"bytes,2,opt,name=read_set,json=readSet" json:"read_set,omitempty"`
	WriteSet     *ConfigGroup      `protobuf:"bytes,3,opt,name=write_set,json=writeSet" json:"write_set,omitempty"`
	IsolatedData map[string][]byte `protobuf:"bytes,5,rep,name=isolated_data,json=isolatedData" json:"isolated_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

type ConfigSignature struct {
	SignatureHeader []byte `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	Signature       []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}
