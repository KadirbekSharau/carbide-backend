package ipfs

// import (
//     "context"
//     "github.com/ipfs/kubo/go-ipfs-files"
//     "github.com/ipfs/kubo/go-ipfs/core/coreapi"
//     "github.com/ipfs/kubo/go-ipfs/core/coreunix"
//     "github.com/ipfs/kubo/go-ipfs/core/node"
//     "github.com/ipfs/kubo/go-ipfs/repo"
//     "io/ioutil"
// )

// func StoreFile(content []byte) (cid *coreapi.ResolvedPath, err error) {
//     // Create a new IPFS node
//     r, err := repo.NewFS(repo.UseMemoryLock(true))
//     if err != nil {
//         return nil, err
//     }
//     ctx, cancel := context.WithCancel(context.Background())
//     defer cancel()
//     nd, err := node.New(ctx, &node.BuildCfg{Repo: r})
//     if err != nil {
//         return nil, err
//     }
//     if err := nd.Start(ctx); err != nil {
//         return nil, err
//     }
//     defer nd.Close()

//     // Create a temporary file with the content
//     file, err := ioutil.TempFile("", "")
//     if err != nil {
//         return nil, err
//     }
//     defer file.Close()
//     _, err = file.Write(content)
//     if err != nil {
//         return nil, err
//     }

//     // Add the file to IPFS
//     ipfsPath, err := coreunix.Add(nd.Context(), nd.Filestore(), files.NewReaderPath(file.Name()))
//     if err != nil {
//         return nil, err
//     }

//     // Resolve the IPFS path to a CID
//     resPath, err := coreapi.ResolvePath(nd.Context(), ipfsPath)
//     if err != nil {
//         return nil, err
//     }

//     return resPath, nil
// }