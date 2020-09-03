// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	job "github.com/jenkins-x/lighthouse/pkg/config/job"
	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityRecord) DeepCopyInto(out *ActivityRecord) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]*ActivityStageOrStep, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ActivityStageOrStep)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]*ActivityStageOrStep, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ActivityStageOrStep)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityRecord.
func (in *ActivityRecord) DeepCopy() *ActivityRecord {
	if in == nil {
		return nil
	}
	out := new(ActivityRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityStageOrStep) DeepCopyInto(out *ActivityStageOrStep) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]*ActivityStageOrStep, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ActivityStageOrStep)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]*ActivityStageOrStep, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ActivityStageOrStep)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityStageOrStep.
func (in *ActivityStageOrStep) DeepCopy() *ActivityStageOrStep {
	if in == nil {
		return nil
	}
	out := new(ActivityStageOrStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ByNum) DeepCopyInto(out *ByNum) {
	{
		in := &in
		*out = make(ByNum, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ByNum.
func (in ByNum) DeepCopy() ByNum {
	if in == nil {
		return nil
	}
	out := new(ByNum)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DecorationConfig) DeepCopyInto(out *DecorationConfig) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(Duration)
		**out = **in
	}
	if in.GracePeriod != nil {
		in, out := &in.GracePeriod, &out.GracePeriod
		*out = new(Duration)
		**out = **in
	}
	if in.SSHKeySecrets != nil {
		in, out := &in.SSHKeySecrets, &out.SSHKeySecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SSHHostFingerprints != nil {
		in, out := &in.SSHHostFingerprints, &out.SSHHostFingerprints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SkipCloning != nil {
		in, out := &in.SkipCloning, &out.SkipCloning
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DecorationConfig.
func (in *DecorationConfig) DeepCopy() *DecorationConfig {
	if in == nil {
		return nil
	}
	out := new(DecorationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Duration) DeepCopyInto(out *Duration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Duration.
func (in *Duration) DeepCopy() *Duration {
	if in == nil {
		return nil
	}
	out := new(Duration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JenkinsSpec) DeepCopyInto(out *JenkinsSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JenkinsSpec.
func (in *JenkinsSpec) DeepCopy() *JenkinsSpec {
	if in == nil {
		return nil
	}
	out := new(JenkinsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LighthouseJob) DeepCopyInto(out *LighthouseJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LighthouseJob.
func (in *LighthouseJob) DeepCopy() *LighthouseJob {
	if in == nil {
		return nil
	}
	out := new(LighthouseJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LighthouseJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LighthouseJobList) DeepCopyInto(out *LighthouseJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LighthouseJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LighthouseJobList.
func (in *LighthouseJobList) DeepCopy() *LighthouseJobList {
	if in == nil {
		return nil
	}
	out := new(LighthouseJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LighthouseJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LighthouseJobSpec) DeepCopyInto(out *LighthouseJobSpec) {
	*out = *in
	if in.Refs != nil {
		in, out := &in.Refs, &out.Refs
		*out = new(Refs)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraRefs != nil {
		in, out := &in.ExtraRefs, &out.ExtraRefs
		*out = make([]Refs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PipelineRunSpec != nil {
		in, out := &in.PipelineRunSpec, &out.PipelineRunSpec
		*out = new(v1beta1.PipelineRunSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PipelineRunParams != nil {
		in, out := &in.PipelineRunParams, &out.PipelineRunParams
		*out = make([]job.PipelineRunParam, len(*in))
		copy(*out, *in)
	}
	if in.PodSpec != nil {
		in, out := &in.PodSpec, &out.PodSpec
		*out = new(v1.PodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JenkinsSpec != nil {
		in, out := &in.JenkinsSpec, &out.JenkinsSpec
		*out = new(JenkinsSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LighthouseJobSpec.
func (in *LighthouseJobSpec) DeepCopy() *LighthouseJobSpec {
	if in == nil {
		return nil
	}
	out := new(LighthouseJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LighthouseJobStatus) DeepCopyInto(out *LighthouseJobStatus) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.Activity != nil {
		in, out := &in.Activity, &out.Activity
		*out = new(ActivityRecord)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LighthouseJobStatus.
func (in *LighthouseJobStatus) DeepCopy() *LighthouseJobStatus {
	if in == nil {
		return nil
	}
	out := new(LighthouseJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pull) DeepCopyInto(out *Pull) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pull.
func (in *Pull) DeepCopy() *Pull {
	if in == nil {
		return nil
	}
	out := new(Pull)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Refs) DeepCopyInto(out *Refs) {
	*out = *in
	if in.Pulls != nil {
		in, out := &in.Pulls, &out.Pulls
		*out = make([]Pull, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Refs.
func (in *Refs) DeepCopy() *Refs {
	if in == nil {
		return nil
	}
	out := new(Refs)
	in.DeepCopyInto(out)
	return out
}
