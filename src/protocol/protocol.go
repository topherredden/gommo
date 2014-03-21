package protocol

const (
    Packet_CreateCharacter byte                 = 0x00
    Packet_Disconnect                      = 0x01
    Packet_MovementReq                     = 0x02
    Packet_AsciiSpeech                     = 0x03
    Packet_GodModeRequest                  = 0x04
    Packet_AttackReq                       = 0x05
    Packet_UseReq                          = 0x06
    Packet_LiftReq                         = 0x07
    Packet_DropReq                         = 0x08
    Packet_LookReq                         = 0x09
    Packet_Edit                            = 0x0A
    Packet_TextCommand                     = 0x12
    Packet_EquipReq                        = 0x13
    Packet_ChangeZ                         = 0x14
    Packet_Resynchronize                   = 0x22
    Packet_DeathStatusResponse             = 0x2C
    Packet_MobileQuery                     = 0x34
    Packet_ChangeSkillLock                 = 0x3A
    Packet_VendorBuyReply                  = 0x3B
    Packet_NewTerrain                      = 0x47
    Packet_NewAnimData                     = 0x48
    Packet_NewRegion                       = 0x58
    Packet_PlayCharacter                   = 0x5D
    Packet_DeleteStatic                    = 0x61
    Packet_TargetResponse                  = 0x6C
    Packet_SecureTrade                     = 0x6F
    Packet_SetWarMode                      = 0x72
    Packet_PingReq                         = 0x73
    Packet_RenameRequest                   = 0x75
    Packet_ResourceQuery                   = 0x79
    Packet_GodviewQuery                    = 0x7E
    Packet_MenuResponse                    = 0x7D
    Packet_AccountLogin byte                    = 0x80
    Packet_DeleteCharacter                 = 0x83
    Packet_GameLogin                       = 0x91
    Packet_HuePickerResponse               = 0x95
    Packet_GameCentralMoniter              = 0x96
    Packet_MobileNameRequest               = 0x98
    Packet_AsciiPromptResponse             = 0x9A
    Packet_HelpRequest                     = 0x9B
    Packet_GMSingle                        = 0x9D
    Packet_VendorSellReply                 = 0x9F
    Packet_PlayServer                      = 0xA0
    Packet_SystemInfo                      = 0xA4
    Packet_RequestScrollWindow             = 0xA7
    Packet_UnicodeSpeech                   = 0xAD
    Packet_DisplayGumpResponse             = 0xB1
    Packet_ChatRequest                     = 0xB5
    Packet_ObjectHelpRequest               = 0xB6
    Packet_ProfileReq                      = 0xB8
    Packet_AccountID                       = 0xBB
    Packet_ClientVersion                   = 0xBD
    Packet_AssistVersion                   = 0xBE
    Packet_ExtendedCommand                 = 0xBF
    Packet_UnicodePromptResponse           = 0xC2
    Packet_SetUpdateRange                  = 0xC8
    Packet_TripTime                        = 0xC9
    Packet_UTripTime                       = 0xCA
    Packet_ConfigurationFile               = 0xD0
    Packet_LogoutReq                       = 0xD1
    Packet_BatchQueryProperties            = 0xD6
    Packet_EncodedCommand                  = 0xD7
    Packet_ClientType                      = 0xE1
    Packet_LoginServerSeed                 = 0xEF
    Packet_CreateCharacter70160            = 0xF8
)
/*
Register6017( 0x08, 15, true, new OnPacketReceive( DropReq6017 ) );

RegisterExtended( 0x05, false, new OnPacketReceive( ScreenSize ) );
RegisterExtended( 0x06,  true, new OnPacketReceive( PartyMessage ) );
RegisterExtended( 0x07,  true, new OnPacketReceive( QuestArrow ) );
RegisterExtended( 0x09,  true, new OnPacketReceive( DisarmRequest ) );
RegisterExtended( 0x0A,  true, new OnPacketReceive( StunRequest ) );
RegisterExtended( 0x0B, false, new OnPacketReceive( Language ) );
RegisterExtended( 0x0C,  true, new OnPacketReceive( CloseStatus ) );
RegisterExtended( 0x0E,  true, new OnPacketReceive( Animate ) );
RegisterExtended( 0x0F, false, new OnPacketReceive( Empty ) ); // What's this?
RegisterExtended( 0x10,  true, new OnPacketReceive( QueryProperties ) );
RegisterExtended( 0x13,  true, new OnPacketReceive( ContextMenuRequest ) );
RegisterExtended( 0x15,  true, new OnPacketReceive( ContextMenuResponse ) );
RegisterExtended( 0x1A,  true, new OnPacketReceive( StatLockChange ) );
RegisterExtended( 0x1C,  true, new OnPacketReceive( CastSpell ) );
RegisterExtended( 0x24, false, new OnPacketReceive( UnhandledBF ) );

RegisterEncoded( 0x19, true, new OnEncodedPacketReceive( SetAbility ) );
RegisterEncoded( 0x28, true, new OnEncodedPacketReceive( GuildGumpRequest ) );

RegisterEncoded( 0x32, true, new OnEncodedPacketReceive( QuestGumpRequest ) );*/
