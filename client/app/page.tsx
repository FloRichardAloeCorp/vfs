"use client";
import {
  Table,
  TableHeader,
  TableBody,
  TableColumn,
  TableRow,
  TableCell,
} from "@nextui-org/table";
import { useEffect, useState } from "react";
import { FolderIcon, DocumentIcon } from "@heroicons/react/24/outline";
import { format } from "date-fns";
import { toast } from "react-toastify";

import { SideMenu } from "@/components/side-menu";
import { usePath } from "@/context/path";
import { File } from "@/types";
import { FilesBreadCrumbs } from "@/components/files-breadcrumbs";
import { FileActionsMenu } from "@/components/file-actions-menu";

export default function Home() {
  const { currentPath, setCurrentPath } = usePath();
  const [currentFiles, setCurrentFiles] = useState<File[]>([]);

  const fetchFiles = async () => {
    try {
      const data = await fetch(
        `${process.env.NEXT_PUBLIC_VFS_BASE_URL}/directory${currentPath}`,
      );
      const files = (await data.json()) as File[];

      setCurrentFiles(files);
    } catch (error) {
      toast("Fail to retrieve files", { type: "error" });
    }
  };

  useEffect(() => {
    fetchFiles();
  }, [currentPath]);

  const handleNewFolderCreated = (newPath: string) => {
    setCurrentPath(newPath);
  };

  const selectFile = (file: File) => {
    return (e: React.MouseEvent) => {
      e.stopPropagation();
      if (file.type == "directory") {
        setCurrentPath(file.path);
      }
    };
  };

  return (
    <div className="flex flex-row h-screen">
      <div className="bg-slate-50 min-w-56">
        <SideMenu onNewFile={fetchFiles} onNewFolder={handleNewFolderCreated} />
      </div>
      <main className="container mx-auto max-w-7xl pt-16 px-6 flex-grow">
        <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
          <FilesBreadCrumbs
            className="place-self-start"
            path={currentPath}
            onItemClicked={setCurrentPath}
          />
          <Table aria-label="Example static collection table">
            <TableHeader>
              <TableColumn key="name">Name</TableColumn>
              <TableColumn key="last_update">Last update</TableColumn>
              <TableColumn key="action">Actions</TableColumn>
            </TableHeader>
            <TableBody emptyContent={"Aucun fichiers à afficher"}>
              {currentFiles.map((file) => (
                <TableRow key={file.id} onClick={selectFile(file)}>
                  <TableCell>
                    <div className="flex flex-row">
                      {file.type == "directory" && (
                        <FolderIcon className="size-5 mr-2" />
                      )}
                      {file.type == "file" && (
                        <DocumentIcon className="size-5 mr-2" />
                      )}
                      {file.name}
                    </div>
                  </TableCell>
                  <TableCell>
                    {format(file.last_update, "dd/MM/yyyy HH:mm")}
                  </TableCell>
                  <TableCell>
                    <FileActionsMenu file={file} onAction={fetchFiles} />
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </section>
      </main>
    </div>
  );
}
