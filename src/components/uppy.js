import React from 'react';
import Uppy from '@uppy/core';
import AwsS3 from '@uppy/aws-s3';
import GoogleDrive from '@uppy/google-drive';
import Url from '@uppy/url';
import { Dashboard } from '@uppy/react';
import { useSelector } from 'react-redux';
import '@uppy/core/dist/style.css';
import '@uppy/dashboard/dist/style.css';
import '@uppy/url/dist/style.css';

function UppyUploader({ onUpload }) {
  // const organisation_slug = useSelector(
  //   (state) => state.organisation.details[state.organisation.selected].slug,
  // );
  const organisation_slug = 'org';

  const uppy = Uppy({
    id: 'uppy-media',
    meta: { type: 'avatar' },
    restrictions: {
      allowedFileTypes: ['.json'],
    },
    autoProceed: false,
    onBeforeUpload: (files) => {
      const updatedFiles = {};
      Object.keys(files).forEach((fileID) => {
        updatedFiles[fileID] = {
          ...files[fileID],
          meta: {
            ...files[fileID].meta,
            name: `bindu
            /${organisation_slug}
            /${new Date().getFullYear()}
            /${new Date().getMonth()}
            /${Date.now().toString()}
            _${files[fileID].meta.name}`,
          },
        };
      });
      return updatedFiles;
    },
  })
    .use(AwsS3, { companionUrl: 'http://127.0.0.1:3020/' })
    .use(Url, { companionUrl: 'http://127.0.0.1:3020/' })
    .use(GoogleDrive, { companionUrl: 'http://127.0.0.11:3020' });

  uppy.on('complete', (result) => {
    const successful = result.successful[0];
    console.log(successful);
    const upload = {};

    upload['file_size'] = successful.size;
    upload['name'] = successful.meta.name;
    upload['type'] = successful.meta.type;
    upload['url'] = {};
    upload['url']['raw'] = successful.uploadURL;

    onUpload(upload);
  });
  return <Dashboard uppy={uppy} plugins={['GoogleDrive', 'Url']} />;
}

export default UppyUploader;
