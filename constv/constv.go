package constv

import "bytes"

const (
	VERSION_FULL    = "v1.0.0.1"
	PACKAGE_NAME    = "spg-core"
	DEAMON_NAME     = "SuperGoldChain"
	DEFAULT_RPCPORT = 32015
)

func GetVersionInfo() string {
	var byteInfo bytes.Buffer
	byteInfo.WriteString("go-spg is the Super Gold Chain written in Golang,")
	byteInfo.WriteString("implements the blockchain hyperledger, supports pow and pos in terms.")
	byteInfo.WriteString("it is the facility for dapps such as wallet,irreverable info databases.")
	byteInfo.WriteString("smart contract is support also,just enjoy and have fun")
	return byteInfo.String()
}

func HelpMessageGroup(message string) string {
	var byteMsg bytes.Buffer
	byteMsg.WriteString(message)
	byteMsg.WriteString("\n\n")
	return byteMsg.String()
}

func HelpMessageOpt(option string, message string) string {
	var byteMsg bytes.Buffer
	byteMsg.WriteString(option)
	byteMsg.WriteString("\n")
	byteMsg.WriteString(message)
	byteMsg.WriteString("\n\n")
	return byteMsg.String()
}
func GetHelpInfo() string {
	// When adding new options to the categories, please keep and ensure alphabetical ordering.
	// Do not translate ...) -help-debug options, Many technical terms, and only a very small audience, so is unnecessary stress to translators.
	var strUsage bytes.Buffer
	strUsage.WriteString(HelpMessageGroup("Options:"))
	strUsage.WriteString(HelpMessageOpt("-?", "Print this help message and exit"))
	strUsage.WriteString(HelpMessageOpt("-version", "Print version and exit"))
	strUsage.WriteString(HelpMessageOpt("-alertnotify=<cmd>", "Execute command when a relevant alert is received or we see a really long fork (%s in cmd is replaced by message)"))
	strUsage.WriteString(HelpMessageOpt("-blocknotify=<cmd>", "Execute command when the best block changes (%s in cmd is replaced by block hash)"))
	//if (showDebug)
	//strUsage.WriteString(HelpMessageOpt("-blocksonly", strprintf("Whether to operate in a blocks only mode (default: %u)", DEFAULT_BLOCKSONLY))
	//strUsage.WriteString(HelpMessageOpt("-assumevalid=<hex>", strprintf("If this block is in the chain assume that it and its ancestors are valid and potentially skip their script verification (0 to verify all, default: %s, testnet: %s)"), Params(CBaseChainParams::MAIN).GetConsensus().defaultAssumeValid.GetHex(), Params(CBaseChainParams::TESTNET).GetConsensus().defaultAssumeValid.GetHex()))
	//strUsage.WriteString(HelpMessageOpt("-conf=<file>", strprintf("Specify configuration file (default: %s)"), BITCOIN_CONF_FILENAME))
	//if (mode == HMM_BITCOIND)
	//{
	//#if HAVE_DECL_DAEMON
	strUsage.WriteString(HelpMessageOpt("-daemon", "Run in the background as a daemon and accept commands"))
	//#endif
	//}
	strUsage.WriteString(HelpMessageOpt("-datadir=<dir>", "Specify data directory"))
	//strUsage.WriteString(HelpMessageOpt("-dbcache=<n>", strprintf("Set database cache size in megabytes (%d to %d, default: %d)"), nMinDbCache, nMaxDbCache, nDefaultDbCache))
	//if (showDebug)
	//strUsage.WriteString(HelpMessageOpt("-feefilter", strprintf("Tell other nodes to filter invs to us by our mempool min fee (default: %u)", DEFAULT_FEEFILTER))
	//strUsage.WriteString(HelpMessageOpt("-loadblock=<file>", "Imports blocks from external blk000??.dat file on startup"))
	//strUsage.WriteString(HelpMessageOpt("-maxorphantx=<n>", strprintf("Keep at most <n> unconnectable transactions in memory (default: %u)"), DEFAULT_MAX_ORPHAN_TRANSACTIONS))
	//strUsage.WriteString(HelpMessageOpt("-maxmempool=<n>", strprintf("Keep the transaction memory pool below <n> megabytes (default: %u)"), DEFAULT_MAX_MEMPOOL_SIZE))
	//strUsage.WriteString(HelpMessageOpt("-mempoolexpiry=<n>", strprintf("Do not keep transactions in the mempool longer than <n> hours (default: %u)"), DEFAULT_MEMPOOL_EXPIRY))
	//strUsage.WriteString(HelpMessageOpt("-blockreconstructionextratxn=<n>", strprintf("Extra transactions to keep in memory for compact block reconstructions (default: %u)"), DEFAULT_BLOCK_RECONSTRUCTION_EXTRA_TXN))
	//strUsage.WriteString(HelpMessageOpt("-par=<n>", strprintf("Set the number of script verification threads (%u to %d, 0 = auto, <0 = leave that many cores free, default: %d)"),
	//-GetNumCores(), MAX_SCRIPTCHECK_THREADS, DEFAULT_SCRIPTCHECK_THREADS))
	////#ifndef WIN32
	//strUsage.WriteString(HelpMessageOpt("-pid=<file>", strprintf("Specify pid file (default: %s)"), BITCOIN_PID_FILENAME))
	////#endif
	//strUsage.WriteString(HelpMessageOpt("-prune=<n>", strprintf("Reduce storage requirements by enabling pruning (deleting) of old blocks. This allows the pruneblockchain RPC to be called to delete specific blocks, and enables automatic pruning of old blocks if a target size in MiB is provided. This mode is incompatible with -txindex and -rescan. "
	//"Warning: Reverting this setting requires re-downloading the entire blockchain. "
	//"(default: 0 = disable pruning blocks, 1 = allow manual pruning via RPC, >%u = automatically prune block files to stay under the specified target size in MiB)"), MIN_DISK_SPACE_FOR_BLOCK_FILES / 1024 / 1024))
	//strUsage.WriteString(HelpMessageOpt("-record-log-opcodes", strprintf("Logs all EVM LOG opcode operations to the file vmExecLogs.json")))
	//strUsage.WriteString(HelpMessageOpt("-reindex-chainstate", "Rebuild chain state from the currently indexed blocks"))
	//strUsage.WriteString(HelpMessageOpt("-reindex", "Rebuild chain state and block index from the blk*.dat files on disk"))
	////#ifndef WIN32
	//strUsage.WriteString(HelpMessageOpt("-sysperms", "Create new files with system default permissions, instead of umask 077 (only effective with disabled wallet functionality)"))
	////#endif
	//strUsage.WriteString(HelpMessageOpt("-txindex", strprintf("Maintain a full transaction index, used by the getrawtransaction rpc call (default: %u)"), DEFAULT_TXINDEX))
	//strUsage.WriteString(HelpMessageOpt("-logevents", strprintf("Maintain a full EVM log index, used by searchlogs and gettransactionreceipt rpc calls (default: %u)"), DEFAULT_LOGEVENTS))
	//
	//strUsage.WriteString(HelpMessageGroup("Connection options:"))
	//strUsage.WriteString(HelpMessageOpt("-addnode=<ip>", "Add a node to connect to and attempt to keep the connection open"))
	//strUsage.WriteString(HelpMessageOpt("-banscore=<n>", strprintf("Threshold for disconnecting misbehaving peers (default: %u)"), DEFAULT_BANSCORE_THRESHOLD))
	//strUsage.WriteString(HelpMessageOpt("-bantime=<n>", strprintf("Number of seconds to keep misbehaving peers from reconnecting (default: %u)"), DEFAULT_MISBEHAVING_BANTIME))
	//strUsage.WriteString(HelpMessageOpt("-bind=<addr>", "Bind to given address and always listen on it. Use [host]:port notation for IPv6"))
	//strUsage.WriteString(HelpMessageOpt("-connect=<ip>", "Connect only to the specified node(s); -noconnect or -connect=0 alone to disable automatic connections"))
	//strUsage.WriteString(HelpMessageOpt("-discover", "Discover own IP addresses (default: 1 when listening and no -externalip or -proxy)"))
	//strUsage.WriteString(HelpMessageOpt("-dns", "Allow DNS lookups for -addnode, -seednode and -connect") + " " + strprintf("(default: %u)"), DEFAULT_NAME_LOOKUP))
	//strUsage.WriteString(HelpMessageOpt("-dnsseed", "Query for peer addresses via DNS lookup, if low on addresses (default: 1 unless -connect/-noconnect)"))
	//strUsage.WriteString(HelpMessageOpt("-externalip=<ip>", "Specify your own public address"))
	//strUsage.WriteString(HelpMessageOpt("-forcednsseed", strprintf("Always query for peer addresses via DNS lookup (default: %u)"), DEFAULT_FORCEDNSSEED))
	//strUsage.WriteString(HelpMessageOpt("-listen", "Accept connections from outside (default: 1 if no -proxy or -connect/-noconnect)"))
	//strUsage.WriteString(HelpMessageOpt("-listenonion", strprintf("Automatically create Tor hidden service (default: %d)"), DEFAULT_LISTEN_ONION))
	//strUsage.WriteString(HelpMessageOpt("-maxconnections=<n>", strprintf("Maintain at most <n> connections to peers (default: %u)"), DEFAULT_MAX_PEER_CONNECTIONS))
	//strUsage.WriteString(HelpMessageOpt("-maxreceivebuffer=<n>", strprintf("Maximum per-connection receive buffer, <n>*1000 bytes (default: %u)"), DEFAULT_MAXRECEIVEBUFFER))
	//strUsage.WriteString(HelpMessageOpt("-maxsendbuffer=<n>", strprintf("Maximum per-connection send buffer, <n>*1000 bytes (default: %u)"), DEFAULT_MAXSENDBUFFER))
	//strUsage.WriteString(HelpMessageOpt("-maxtimeadjustment", strprintf("Maximum allowed median peer time offset adjustment. Local perspective of time may be influenced by peers forward or backward by this amount. (default: %u seconds)"), DEFAULT_MAX_TIME_ADJUSTMENT))
	//strUsage.WriteString(HelpMessageOpt("-onion=<ip:port>", strprintf("Use separate SOCKS5 proxy to reach peers via Tor hidden services (default: %s)"), "-proxy"))
	//strUsage.WriteString(HelpMessageOpt("-onlynet=<net>", "Only connect to nodes in network <net> (ipv4, ipv6 or onion)"))
	//strUsage.WriteString(HelpMessageOpt("-permitbaremultisig", strprintf("Relay non-P2SH multisig (default: %u)"), DEFAULT_PERMIT_BAREMULTISIG))
	//strUsage.WriteString(HelpMessageOpt("-peerbloomfilters", strprintf("Support filtering of blocks and transaction with bloom filters (default: %u)"), DEFAULT_PEERBLOOMFILTERS))
	//strUsage.WriteString(HelpMessageOpt("-port=<port>", strprintf("Listen for connections on <port> (default: %u or testnet: %u)"), Params(CBaseChainParams::MAIN).GetDefaultPort(), Params(CBaseChainParams::TESTNET).GetDefaultPort()))
	//strUsage.WriteString(HelpMessageOpt("-proxy=<ip:port>", "Connect through SOCKS5 proxy"))
	//strUsage.WriteString(HelpMessageOpt("-proxyrandomize", strprintf("Randomize credentials for every proxy connection. This enables Tor stream isolation (default: %u)"), DEFAULT_PROXYRANDOMIZE))
	//strUsage.WriteString(HelpMessageOpt("-rpcserialversion", strprintf("Sets the serialization of raw transaction or block hex returned in non-verbose mode, non-segwit(0) or segwit(1) (default: %d)"), DEFAULT_RPC_SERIALIZE_VERSION))
	//strUsage.WriteString(HelpMessageOpt("-seednode=<ip>", "Connect to a node to retrieve peer addresses, and disconnect"))
	//strUsage.WriteString(HelpMessageOpt("-timeout=<n>", strprintf("Specify connection timeout in milliseconds (minimum: 1, default: %d)"), DEFAULT_CONNECT_TIMEOUT))
	//strUsage.WriteString(HelpMessageOpt("-torcontrol=<ip>:<port>", strprintf("Tor control port to use if onion listening enabled (default: %s)"), DEFAULT_TOR_CONTROL))
	//strUsage.WriteString(HelpMessageOpt("-torpassword=<pass>", "Tor control port password (default: empty)"))
	//strUsage.WriteString(HelpMessageOpt("-dgpstorage", "Receiving data from DGP via storage (default: -dgpevm)"))
	//strUsage.WriteString(HelpMessageOpt("-dgpevm", "Receiving data from DGP via a contract call (default: -dgpevm)"))
	////#ifdef USE_UPNP
	////#if USE_UPNP
	//strUsage.WriteString(HelpMessageOpt("-upnp", "Use UPnP to map the listening port (default: 1 when listening and no -proxy)"))
	////#else
	//strUsage.WriteString(HelpMessageOpt("-upnp", strprintf("Use UPnP to map the listening port (default: %u)"), 0))
	////#endif
	////#endif
	//strUsage.WriteString(HelpMessageOpt("-whitebind=<addr>", "Bind to given address and whitelist peers connecting to it. Use [host]:port notation for IPv6"))
	//strUsage.WriteString(HelpMessageOpt("-whitelist=<IP address or network>", "Whitelist peers connecting from the given IP address (e.g. 1.2.3.4) or CIDR notated network (e.g. 1.2.3.0/24). Can be specified multiple times.") +
	//" " + "Whitelisted peers cannot be DoS banned and their transactions are always relayed, even if they are already in the mempool, useful e.g. for a gateway"))
	//strUsage.WriteString(HelpMessageOpt("-whitelistrelay", strprintf("Accept relayed transactions received from whitelisted peers even when not relaying transactions (default: %d)"), DEFAULT_WHITELISTRELAY))
	//strUsage.WriteString(HelpMessageOpt("-whitelistforcerelay", strprintf("Force relay of transactions from whitelisted peers even if they violate local relay policy (default: %d)"), DEFAULT_WHITELISTFORCERELAY))
	//strUsage.WriteString(HelpMessageOpt("-maxuploadtarget=<n>", strprintf("Tries to keep outbound traffic under the given target (in MiB per 24h), 0 = no limit (default: %d)"), DEFAULT_MAX_UPLOAD_TARGET))
	//
	////#ifdef ENABLE_WALLET
	//strUsage.WriteString(CWallet::GetWalletHelpString(showDebug);
	////#endif
	//
	////#if ENABLE_ZMQ
	//strUsage.WriteString(HelpMessageGroup("ZeroMQ notification options:"))
	//strUsage.WriteString(HelpMessageOpt("-zmqpubhashblock=<address>", "Enable publish hash block in <address>"))
	//strUsage.WriteString(HelpMessageOpt("-zmqpubhashtx=<address>", "Enable publish hash transaction in <address>"))
	//strUsage.WriteString(HelpMessageOpt("-zmqpubrawblock=<address>", "Enable publish raw block in <address>"))
	//strUsage.WriteString(HelpMessageOpt("-zmqpubrawtx=<address>", "Enable publish raw transaction in <address>"))
	////#endif
	//
	//strUsage.WriteString(HelpMessageGroup("Debugging/Testing options:"))
	//strUsage.WriteString(HelpMessageOpt("-uacomment=<cmt>", "Append comment to the user agent string"))
	////if (showDebug)
	////{
	//strUsage.WriteString(HelpMessageOpt("-checkblocks=<n>", strprintf("How many blocks to check at startup (default: %u, 0 = all)"), DEFAULT_CHECKBLOCKS))
	//strUsage.WriteString(HelpMessageOpt("-checklevel=<n>", strprintf("How thorough the block verification of -checkblocks is (0-4, default: %u)"), DEFAULT_CHECKLEVEL))
	//strUsage.WriteString(HelpMessageOpt("-checkblockindex", strprintf("Do a full consistency check for mapBlockIndex, setBlockIndexCandidates, chainActive and mapBlocksUnlinked occasionally. Also sets -checkmempool (default: %u)", Params(CBaseChainParams::MAIN).DefaultConsistencyChecks()))
	//strUsage.WriteString(HelpMessageOpt("-checkmempool=<n>", strprintf("Run checks every <n> transactions (default: %u)", Params(CBaseChainParams::MAIN).DefaultConsistencyChecks()))
	//strUsage.WriteString(HelpMessageOpt("-checkpoints", strprintf("Disable expensive verification for known chain history (default: %u)", DEFAULT_CHECKPOINTS_ENABLED))
	//strUsage.WriteString(HelpMessageOpt("-disablesafemode", strprintf("Disable safemode, override a real safe mode event (default: %u)", DEFAULT_DISABLE_SAFEMODE))
	//strUsage.WriteString(HelpMessageOpt("-testsafemode", strprintf("Force safe mode (default: %u)", DEFAULT_TESTSAFEMODE))
	//strUsage.WriteString(HelpMessageOpt("-dropmessagestest=<n>", "Randomly drop 1 of every <n> network messages");
	//strUsage.WriteString(HelpMessageOpt("-fuzzmessagestest=<n>", "Randomly fuzz 1 of every <n> network messages");
	//strUsage.WriteString(HelpMessageOpt("-stopafterblockimport", strprintf("Stop running after importing blocks from disk (default: %u)", DEFAULT_STOPAFTERBLOCKIMPORT))
	//strUsage.WriteString(HelpMessageOpt("-limitancestorcount=<n>", strprintf("Do not accept transactions if number of in-mempool ancestors is <n> or more (default: %u)", DEFAULT_ANCESTOR_LIMIT))
	//strUsage.WriteString(HelpMessageOpt("-limitancestorsize=<n>", strprintf("Do not accept transactions whose size with all in-mempool ancestors exceeds <n> kilobytes (default: %u)", DEFAULT_ANCESTOR_SIZE_LIMIT))
	//strUsage.WriteString(HelpMessageOpt("-limitdescendantcount=<n>", strprintf("Do not accept transactions if any ancestor would have <n> or more in-mempool descendants (default: %u)", DEFAULT_DESCENDANT_LIMIT))
	//strUsage.WriteString(HelpMessageOpt("-limitdescendantsize=<n>", strprintf("Do not accept transactions if any ancestor would have more than <n> kilobytes of in-mempool descendants (default: %u).", DEFAULT_DESCENDANT_SIZE_LIMIT))
	//strUsage.WriteString(HelpMessageOpt("-bip9params=deployment:start:end", "Use given start/end times for specified BIP9 deployment (regtest-only)");
	////}
	//std::string debugCategories = "addrman, alert, bench, cmpctblock, coindb, db, http, libevent, lock, mempool, mempoolrej, miner, net, proxy, prune, rand, reindex, rpc, selectcoins, tor, zmq"; // Don't translate these and qt below
	////if (mode == HMM_BITCOIN_QT)
	//debugCategories += ", qt";
	//strUsage.WriteString(HelpMessageOpt("-debug=<category>", strprintf("Output debugging information (default: %u, supplying <category> is optional)"), 0) + ". " +
	//("If <category> is not supplied or if <category> = 1, output all debugging information.") + ("<category> can be:") + " " + debugCategories + ".");
	////if (showDebug)
	//strUsage.WriteString(HelpMessageOpt("-nodebug", "Turn off debugging messages, same as -debug=0");
	//strUsage.WriteString(HelpMessageOpt("-help-debug", "Show all debugging options (usage: --help -help-debug)"))
	//strUsage.WriteString(HelpMessageOpt("-logips", strprintf("Include IP addresses in debug output (default: %u)"), DEFAULT_LOGIPS))
	//strUsage.WriteString(HelpMessageOpt("-logtimestamps", strprintf("Prepend debug output with timestamp (default: %u)"), DEFAULT_LOGTIMESTAMPS))
	////if (showDebug)
	////{
	//strUsage.WriteString(HelpMessageOpt("-logtimemicros", strprintf("Add microsecond precision to debug timestamps (default: %u)", DEFAULT_LOGTIMEMICROS))
	//strUsage.WriteString(HelpMessageOpt("-mocktime=<n>", "Replace actual time with <n> seconds since epoch (default: 0)");
	//strUsage.WriteString(HelpMessageOpt("-limitfreerelay=<n>", strprintf("Continuously rate-limit free transactions to <n>*1000 bytes per minute (default: %u)", DEFAULT_LIMITFREERELAY))
	//strUsage.WriteString(HelpMessageOpt("-relaypriority", strprintf("Require high priority for relaying free or low-fee transactions (default: %u)", DEFAULT_RELAYPRIORITY))
	//strUsage.WriteString(HelpMessageOpt("-maxsigcachesize=<n>", strprintf("Limit size of signature cache to <n> MiB (default: %u)", DEFAULT_MAX_SIG_CACHE_SIZE))
	//strUsage.WriteString(HelpMessageOpt("-maxtipage=<n>", strprintf("Maximum tip age in seconds to consider node in initial block download (default: %u)", DEFAULT_MAX_TIP_AGE))
	////}
	//strUsage.WriteString(HelpMessageOpt("-minmempoolgaslimit=<limit>", strprintf("The minimum transaction gas limit we are willing to accept into the mempool (default: %s)",MEMPOOL_MIN_GAS_LIMIT))
	//strUsage.WriteString(HelpMessageOpt("-minrelaytxfee=<amt>", strprintf("Fees (in %s/kB) smaller than this are considered zero fee for relaying, mining and transaction creation (default: %s)"),
	//CURRENCY_UNIT, FormatMoney(DEFAULT_MIN_RELAY_TX_FEE)))
	//strUsage.WriteString(HelpMessageOpt("-maxtxfee=<amt>", strprintf("Maximum total fees (in %s) to use in a single wallet transaction or raw transaction; setting this too low may abort large transactions (default: %s)"),
	//CURRENCY_UNIT, FormatMoney(DEFAULT_TRANSACTION_MAXFEE)))
	//strUsage.WriteString(HelpMessageOpt("-printtoconsole", "Send trace/debug info to console instead of debug.log file"))
	////if (showDebug)
	////{
	//strUsage.WriteString(HelpMessageOpt("-printpriority", strprintf("Log transaction priority and fee per kB when mining blocks (default: %u)", DEFAULT_PRINTPRIORITY))
	////}
	//strUsage.WriteString(HelpMessageOpt("-shrinkdebugfile", "Shrink debug.log file on client startup (default: 1 when no -debug)"))
	//
	//AppendParamsHelpMessages(strUsage, showDebug);
	//
	//strUsage.WriteString(HelpMessageGroup("Node relay options:"))
	////if (showDebug) {
	//strUsage.WriteString(HelpMessageOpt("-acceptnonstdtxn", strprintf("Relay and mine \"non-standard\" transactions (%sdefault: %u)", "testnet/regtest only; ", !Params(CBaseChainParams::TESTNET).RequireStandard()))
	//strUsage.WriteString(HelpMessageOpt("-incrementalrelayfee=<amt>", strprintf("Fee rate (in %s/kB) used to define cost of relay, used for mempool limiting and BIP 125 replacement. (default: %s)", CURRENCY_UNIT, FormatMoney(DEFAULT_INCREMENTAL_RELAY_FEE)))
	//strUsage.WriteString(HelpMessageOpt("-dustrelayfee=<amt>", strprintf("Fee rate (in %s/kB) used to defined dust, the value of an output such that it will cost about 1/3 of its value in fees at this fee rate to spend it. (default: %s)", CURRENCY_UNIT, FormatMoney(DUST_RELAY_TX_FEE)))
	////}
	//strUsage.WriteString(HelpMessageOpt("-bytespersigop", strprintf("Equivalent bytes per sigop in transactions for relay and mining (default: %u)"), DEFAULT_BYTES_PER_SIGOP))
	//strUsage.WriteString(HelpMessageOpt("-datacarrier", strprintf("Relay and mine data carrier transactions (default: %u)"), DEFAULT_ACCEPT_DATACARRIER))
	//strUsage.WriteString(HelpMessageOpt("-datacarriersize", strprintf("Maximum size of data in data carrier transactions we relay and mine (default: %u)"), MAX_OP_RETURN_RELAY))
	//strUsage.WriteString(HelpMessageOpt("-mempoolreplacement", strprintf("Enable transaction replacement in the memory pool (default: %u)"), DEFAULT_ENABLE_REPLACEMENT))
	//
	//strUsage.WriteString(HelpMessageGroup("Block creation options:"))
	//strUsage.WriteString(HelpMessageOpt("-blockmaxweight=<n>", strprintf("Set maximum BIP141 block weight (default: %d)"), DEFAULT_BLOCK_MAX_WEIGHT))
	//strUsage.WriteString(HelpMessageOpt("-blockmaxsize=<n>", strprintf("Set maximum block size in bytes (default: %d)"), DEFAULT_BLOCK_MAX_SIZE))
	//strUsage.WriteString(HelpMessageOpt("-blockprioritysize=<n>", strprintf("Set maximum size of high-priority/low-fee transactions in bytes (default: %d)"), DEFAULT_BLOCK_PRIORITY_SIZE))
	//strUsage.WriteString(HelpMessageOpt("-blockmintxfee=<amt>", strprintf("Set lowest fee rate (in %s/kB) for transactions to be included in block creation. (default: %s)"), CURRENCY_UNIT, FormatMoney(DEFAULT_BLOCK_MIN_TX_FEE)))
	//
	//strUsage.WriteString(HelpMessageOpt("-staker-min-tx-gas-price=<amt>", "Any contract execution with a gas price below this will not be included in a block (defaults to the value specified by the DGP)"))
	//strUsage.WriteString(HelpMessageOpt("-staker-max-tx-gas-limit=<n>", "Any contract execution with a gas limit over this amount will not be included in a block (defaults to soft block gas limit)"))
	//strUsage.WriteString(HelpMessageOpt("-staker-soft-block-gas-limit=<n>", "After this amount of gas is surpassed in a block, no more contract executions will be added to the block (defaults to consensus-critical maximum block gas limit)"))
	//
	////if (showDebug)
	//strUsage.WriteString(HelpMessageOpt("-blockversion=<n>", "Override block version to test forking scenarios");
	//
	//strUsage.WriteString(HelpMessageGroup("RPC server options:"))
	//strUsage.WriteString(HelpMessageOpt("-server", "Accept command line and JSON-RPC commands"))
	//strUsage.WriteString(HelpMessageOpt("-rest", strprintf("Accept public REST requests (default: %u)"), DEFAULT_REST_ENABLE))
	//strUsage.WriteString(HelpMessageOpt("-rpcbind=<addr>", "Bind to given address to listen for JSON-RPC connections. Use [host]:port notation for IPv6. This option can be specified multiple times (default: bind to all interfaces)"))
	//strUsage.WriteString(HelpMessageOpt("-rpccookiefile=<loc>", "Location of the auth cookie (default: data dir)"))
	//strUsage.WriteString(HelpMessageOpt("-rpcuser=<user>", "Username for JSON-RPC connections"))
	//strUsage.WriteString(HelpMessageOpt("-rpcpassword=<pw>", "Password for JSON-RPC connections"))
	//strUsage.WriteString(HelpMessageOpt("-rpcauth=<userpw>", "Username and hashed password for JSON-RPC connections. The field <userpw> comes in the format: <USERNAME>:<SALT>$<HASH>. A canonical python script is included in share/rpcuser. The client then connects normally using the rpcuser=<USERNAME>/rpcpassword=<PASSWORD> pair of arguments. This option can be specified multiple times"))
	//strUsage.WriteString(HelpMessageOpt("-rpcport=<port>", strprintf("Listen for JSON-RPC connections on <port> (default: %u or testnet: %u)"), BaseParams(CBaseChainParams::MAIN).RPCPort(), BaseParams(CBaseChainParams::TESTNET).RPCPort()))
	//strUsage.WriteString(HelpMessageOpt("-rpcallowip=<ip>", "Allow JSON-RPC connections from specified source. Valid for <ip> are a single IP (e.g. 1.2.3.4), a network/netmask (e.g. 1.2.3.4/255.255.255.0) or a network/CIDR (e.g. 1.2.3.4/24). This option can be specified multiple times"))
	//strUsage.WriteString(HelpMessageOpt("-rpcthreads=<n>", strprintf("Set the number of threads to service RPC calls (default: %d)"), DEFAULT_HTTP_THREADS))
	////if (showDebug) {
	//strUsage.WriteString(HelpMessageOpt("-rpcworkqueue=<n>", strprintf("Set the depth of the work queue to service RPC calls (default: %d)", DEFAULT_HTTP_WORKQUEUE))
	//strUsage.WriteString(HelpMessageOpt("-rpcservertimeout=<n>", strprintf("Timeout during HTTP requests (default: %d)", DEFAULT_HTTP_SERVER_TIMEOUT))
	//}

	return strUsage.String()
}
