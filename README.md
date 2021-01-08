# GDB 
The GDB database is a real-time database written in Go language based on LevelDb. It can add real-time data in real time and obtain the history of the data. It is suitable for a large number of
The process of storing data such as thermal process, real-time web application, etc.
The structure of the entire database is divided into groups Group and Point Item. The group is equivalent to the table in SQL, and the Item is equivalent to the data in each table. Before use
To add the corresponding Group first, then add Item to the Group. After having the Item, you can write real-time data to the Item, after writing the real-time data
You can get the corresponding historical data.
Features of GDB:
1. Based on LevelDB, batch write speed is very fast, supporting thousands of connections to write concurrently at the same time
2. Using bloomFilter, batch reading is also fast
3. Provides a high-performance, concurrent webRestful interface, provides web programs and client programs, and provides go language users
The underlying functions are easy to use and can be seamlessly connected to various languages.
3. Suitable for occasions that require large-scale storage of data, such as the storage of thermal process data in power plants, and the storage of real-time web-related data
The data in the Group is stored in SQLite, and the real-time data and historical data are stored in Leveldb. The structure of the entire database is as follows:
4. A web application is provided, which can conveniently operate the GDB database
5. Provides secondary calculation based on js, enabling secondary development of real-time data
6. see document for details:  https://justkeepsilence.gitbook.io/gdb/